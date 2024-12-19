package repository_test

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/gameserver"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
)

var _ = Describe("Connection repository", func() {
	var repo repository.ConnectionRepository
	BeforeEach(func() {
		repo = repository.NewPgxConnectionRepository(migrater)
		Expect(repo).NotTo(BeNil())
	})

	Describe("CreatePendingConnection", func() {
		It("should create a pending connection", func(ctx SpecContext) {
			characterId := uuid.New()
			serverName := faker.Username()
			conn, err := repo.CreatePendingConnection(ctx, &characterId, serverName)
			Expect(err).NotTo(HaveOccurred())
			Expect(conn).NotTo(BeNil())
			Expect(conn.CharacterId).To(Equal(characterId))
			Expect(conn.ServerName).To(Equal(serverName))
			Expect(conn.Id).NotTo(BeNil())
		})
		It("should error given an empty server name", func(ctx SpecContext) {
			characterId := uuid.New()
			conn, err := repo.CreatePendingConnection(ctx, &characterId, "")
			Expect(err).To(HaveOccurred())
			Expect(conn).To(BeNil())
		})
		It("should error given a nil character", func(ctx SpecContext) {
			conn, err := repo.CreatePendingConnection(ctx, nil, faker.Username())
			Expect(err).To(HaveOccurred())
			Expect(conn).To(BeNil())
		})
	})

	Context("when a pending connection exists", func() {
		var conn *gameserver.PendingConnection
		BeforeEach(func(ctx SpecContext) {
			var err error
			characterId := uuid.New()
			serverName := faker.Username()
			conn, err = repo.CreatePendingConnection(ctx, &characterId, serverName)
			Expect(err).NotTo(HaveOccurred())
			Expect(conn).NotTo(BeNil())
			Expect(conn.CharacterId).To(Equal(characterId))
			Expect(conn.ServerName).To(Equal(serverName))
			Expect(conn.Id).NotTo(BeNil())
		})
		Describe("DeletePendingConnection", func() {
			It("should delete a pending connection", func(ctx SpecContext) {
				Expect(repo.DeletePendingConnection(ctx, &conn.Id)).To(Succeed())
			})
			It("should error not found given a random id", func(ctx SpecContext) {
				id := uuid.New()
				Expect(repo.DeletePendingConnection(ctx, &id)).To(Equal(pgx.ErrNoRows))
			})
		})
		Describe("FindPendingConnection", func() {
			BeforeEach(func(ctx SpecContext) {
				pc, err := repo.FindPendingConnection(ctx, &conn.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(pc).NotTo(BeNil())
				Expect(pc.CharacterId).To(Equal(conn.CharacterId))
				Expect(pc.ServerName).To(Equal(conn.ServerName))
				Expect(pc.Id).To(Equal(conn.Id))
			})

			It("should find a pending connection", func(ctx SpecContext) {
			})

			It("should return error not found given a random id", func(ctx SpecContext) {
				id := uuid.New()
				pc, err := repo.FindPendingConnection(ctx, &id)
				Expect(err).To(Equal(pgx.ErrNoRows))
				Expect(pc).To(BeNil())
			})

			It("should not find pending connections if deleted", func(ctx SpecContext) {
				Expect(repo.DeletePendingConnection(ctx, &conn.Id)).To(Succeed())
				pc, err := repo.FindPendingConnection(ctx, &conn.Id)
				Expect(err).To(Equal(pgx.ErrNoRows))
				Expect(pc).To(BeNil())
			})
		})
	})
})
