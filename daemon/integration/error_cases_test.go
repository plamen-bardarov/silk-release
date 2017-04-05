package integration_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"code.cloudfoundry.org/go-db-helpers/testsupport"
	"code.cloudfoundry.org/silk/daemon/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("error cases", func() {
	var (
		testDatabase *testsupport.TestDatabase
		daemonConfig config.Config
	)

	BeforeEach(func() {
		dbName := fmt.Sprintf("test_database_%x", GinkgoParallelNode())
		dbConnectionInfo := testsupport.GetDBConnectionInfo()
		testDatabase = dbConnectionInfo.CreateDatabase(dbName)

		daemonConfig = config.Config{
			SubnetRange: "10.255.0.0/16",
			SubnetMask:  24,
			Database:    testDatabase.DBConfig(),
			UnderlayIP:  "10.244.4.6",
		}
	})

	AfterEach(func() {
		if testDatabase != nil {
			testDatabase.Destroy()
		}
	})

	Context("when the path to the config is bad", func() {
		It("exits with status 1", func() {
			startCmd := exec.Command(daemonPath, "--config", "some/bad/path")
			session, err := gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("could not read config file"))

			session.Interrupt()
		})
	})

	Context("when the contents of the config file cannot be unmarshaled", func() {
		It("exits with status 1", func() {
			configFile, err := ioutil.TempFile("", "test-config")
			Expect(err).NotTo(HaveOccurred())

			err = ioutil.WriteFile(configFile.Name(), []byte("some-bad-contents"), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			startCmd := exec.Command(daemonPath, "--config", configFile.Name())
			session, err := gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("could not unmarshal config file contents"))

			session.Interrupt()
		})
	})

	Context("when the config has an unsupported type", func() {
		It("exits with status 1", func() {
			daemonConfig.Database.Type = "bad-type"
			configFilePath := writeConfigFile(daemonConfig)

			startCmd := exec.Command(daemonPath, "--config", configFilePath)
			session, err := gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("could not connect to database:"))

			session.Interrupt()
		})
	})

	Context("when the config has a bad connection string", func() {
		It("exits with status 1", func() {
			daemonConfig.Database.ConnectionString = "some-bad-connection-string"

			configFilePath := writeConfigFile(daemonConfig)

			startCmd := exec.Command(daemonPath, "--config", configFilePath)
			session, err := gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, 10*time.Second).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("could not connect to database:"))

			session.Interrupt()
		})
	})

	// TODO(gabe): unpend, figure out how to set up the test so that we can trigger
	// this sort of failure and actually test the behavior in that case
	XContext("when the lease controller fails to acquire a subnet lease", func() {
		It("exits with status 1", func() {
			conf := config.Config{
				SubnetRange: "10.255.0.0/16",
				SubnetMask:  24,
				Database:    testDatabase.DBConfig(),
				UnderlayIP:  "10.244.4.5",
			}

			configFilePath := writeConfigFile(conf)
			startCmd := exec.Command(daemonPath, "--config", configFilePath)
			session, err := gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session.Out, "4s").Should(gbytes.Say("subnet-acquired.*subnet.*underlay ip.*"))

			failingCmd := exec.Command(daemonPath, "--config", configFilePath)
			failingSession, err := gexec.Start(failingCmd, GinkgoWriter, GinkgoWriter)
			Eventually(failingSession, 20*time.Second).Should(gexec.Exit(1))
			Expect(failingSession.Err.Contents()).To(ContainSubstring("could not acquire subnet:"))

			session.Interrupt()
			failingSession.Interrupt()
		})
	})
})
