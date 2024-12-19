package repository_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
)

var _ = Describe("Connection repository", func() {
	var repo repository.ConnectionRepository
	BeforeEach(func() {
		repo = repository.NewPgxConnectionRepository(migrater)
		Expect(repo).NotTo(BeNil())
	})

	Describe("CreatePendingConnection", func() {
		It("should create a pending connection", func() {
			characterName := faker.Username()
			serverName := faker.Username()
			conn, err := repo.CreatePendingConnection(ctx, characterName, serverName)
			Expect(err).NotTo(HaveOccurred())
			Expect(conn).NotTo(BeNil())
			Expect(conn.Character).To(Equal(characterName))
			Expect(conn.ServerName).To(Equal(serverName))
			Expect(conn.Id).NotTo(BeNil())
		})
	})
})
