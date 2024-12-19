package repository_test

import (
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
)

var _ = Describe("Map", func() {
	var repo repository.MapRepository
	var m *game.Map
	BeforeEach(func() {
		repo = repository.NewPgxMapRepository(migrater)
		Expect(repo).NotTo(BeNil())
		m = &game.Map{
			Name:    faker.Username(),
			MapPath: faker.Username(),
		}
	})

	Describe("CreateMap", func() {
		It("should create a map", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(outM, m)
		})
		When("given invalid input", func() {
			var err error
			var outM *game.Map
			AfterEach(func() {
				Expect(err).To(HaveOccurred())
				Expect(outM).To(BeNil())
			})
			It("should error given a nil map", func(ctx SpecContext) {
				m = nil
				outM, err = repo.CreateMap(ctx, nil)
			})
			It("should error given an empty name", func(ctx SpecContext) {
				m.Name = ""
				outM, err = repo.CreateMap(ctx, m)
			})
			It("should error given an empty map path", func(ctx SpecContext) {
				m.MapPath = ""
				outM, err = repo.CreateMap(ctx, m)
			})
			It("should error given an existing map name", func(ctx SpecContext) {
				outM, err = repo.CreateMap(ctx, m)
				Expect(err).NotTo(HaveOccurred())
				Expect(outM).NotTo(BeNil())
				outM, err = repo.CreateMap(ctx, m)
			})
		})
	})

	Describe("DeleteMap", func() {
		It("should delete a map", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			outM, err = repo.DeleteMap(ctx, &outM.Id)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(outM, m)
		})
		It("should delete a dimension map association", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())

			maps := game.Maps{outM}

			m.Name = faker.Username()
			outM, err = repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			maps = append(maps, outM)

			dRepo := repository.NewPgxDimensionRepository(migrater)
			Expect(dRepo).NotTo(BeNil())
			dim, err := dRepo.CreateDimension(ctx, &game.Dimension{
				Name:     faker.Username(),
				Location: faker.Username(),
				Version:  faker.Username(),
				Maps:     maps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(dim).NotTo(BeNil())
			Expect(dim.Maps).To(HaveLen(2))
			Expect(dim.Maps).To(ContainElement(outM))

			outM, err = repo.DeleteMap(ctx, &outM.Id)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(outM, m)

			dim, err = dRepo.GetDimensionById(ctx, &dim.Id)
			Expect(err).NotTo(HaveOccurred())
			Expect(dim).NotTo(BeNil())
			Expect(dim.Maps).To(HaveLen(1))
		})
		When("given invalid input", func() {
			var err error
			var outM *game.Map
			AfterEach(func() {
				Expect(err).To(HaveOccurred())
				Expect(outM).To(BeNil())
			})
			It("should error given a nil map id", func(ctx SpecContext) {
				outM, err = repo.DeleteMap(ctx, nil)
			})
			It("should error given a non-existent map id", func(ctx SpecContext) {
				id := uuid.New()
				outM, err = repo.DeleteMap(ctx, &id)
			})
		})
	})

	Describe("GetDeletedMaps", func() {
		It("should get deleted maps", func(ctx SpecContext) {
			for i := 0; i < 5; i++ {
				m.Name = faker.Username()
				m.MapPath = faker.Username()
				outM, err := repo.CreateMap(ctx, m)
				Expect(err).NotTo(HaveOccurred())
				Expect(outM).NotTo(BeNil())
				outM, err = repo.DeleteMap(ctx, &outM.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(outM).NotTo(BeNil())
				maps, err := repo.GetDeletedMaps(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(maps)).To(BeNumerically(">", i))
				Expect(maps).To(ContainElement(outM))
			}
		})
	})

	Describe("GetMapById", func() {
		It("should get a map by id", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			newM, err := repo.GetMapById(ctx, &outM.Id)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(newM, outM)
		})
		It("should get a deleted map by id", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			outM, err = repo.DeleteMap(ctx, &outM.Id)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			newM, err := repo.GetMapById(ctx, &outM.Id)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(newM, outM)
		})
		When("given invalid input", func() {
			var err error
			var outM *game.Map
			AfterEach(func() {
				Expect(err).To(HaveOccurred())
				Expect(outM).To(BeNil())
			})
			It("should error given a nil map id", func(ctx SpecContext) {
				outM, err = repo.GetMapById(ctx, nil)
			})
			It("should error given a non-existent map id", func(ctx SpecContext) {
				id := uuid.New()
				outM, err = repo.GetMapById(ctx, &id)
			})
		})
	})

	Describe("GetMaps", func() {
		It("should get maps", func(ctx SpecContext) {
			for i := 0; i < 5; i++ {
				m.Name = faker.Username()
				m.MapPath = faker.Username()
				outM, err := repo.CreateMap(ctx, m)
				Expect(err).NotTo(HaveOccurred())
				Expect(outM).NotTo(BeNil())
				maps, err := repo.GetMaps(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(maps)).To(BeNumerically(">", i))
				Expect(maps).To(ContainElement(outM))
				for _, m := range maps {
					Expect(m.DeletedAt).To(BeNil())
				}
			}
		})
		It("should not get deleted maps", func(ctx SpecContext) {
			for i := 0; i < 5; i++ {
				m.Name = faker.Username()
				m.MapPath = faker.Username()
				outM, err := repo.CreateMap(ctx, m)
				Expect(err).NotTo(HaveOccurred())
				Expect(outM).NotTo(BeNil())
				deletedM, err := repo.DeleteMap(ctx, &outM.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(outM).NotTo(BeNil())
				expectEqual(deletedM, outM)
				maps, err := repo.GetMaps(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(maps).NotTo(ContainElement(outM))
			}
		})
	})

	Describe("UpdateMap", func() {
		It("should update a map", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			outM.Name = faker.Username()
			outM.MapPath = faker.Username()
			newM, err := repo.UpdateMap(ctx, outM)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(newM, outM)
		})
		It("should update a map with a new name", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())
			outM.Name = faker.Username()
			newM, err := repo.UpdateMap(ctx, outM)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(newM, outM)
		})
		It("should update a map with new values", func(ctx SpecContext) {
			outM, err := repo.CreateMap(ctx, m)
			Expect(err).NotTo(HaveOccurred())
			Expect(outM).NotTo(BeNil())

			deletedAt := time.Now()
			outM.MapPath = faker.Username()
			outM.Name = faker.Username()
			outM.CreatedAt = time.Now().Add(-time.Hour)
			outM.UpdatedAt = time.Now().Add(-time.Hour)
			outM.DeletedAt = &deletedAt

			newM, err := repo.UpdateMap(ctx, outM)
			Expect(err).NotTo(HaveOccurred())
			expectEqual(newM, outM)
			Expect(newM.UpdatedAt).To(BeTemporally(">", outM.UpdatedAt), "updated at should be greater than the original")
			Expect(newM.DeletedAt).To(BeNil(), "deleted at should not change")
		})

		When("given invalid input", func() {
			var err error
			var outM *game.Map
			AfterEach(func() {
				Expect(err).To(HaveOccurred())
				Expect(outM).To(BeNil())
			})
			It("should error given a nil map", func(ctx SpecContext) {
				outM, err = repo.UpdateMap(ctx, nil)
			})
			It("should error given an empty name", func(ctx SpecContext) {
				m.Name = ""
				outM, err = repo.UpdateMap(ctx, m)
			})
			It("should error given an empty map path", func(ctx SpecContext) {
				m.MapPath = ""
				outM, err = repo.UpdateMap(ctx, m)
			})
			It("should error given a non-existent map id", func(ctx SpecContext) {
				m.Id = uuid.New()
				outM, err = repo.UpdateMap(ctx, m)
			})
		})
	})
})

func expectEqual(new, original *game.Map) {
	Expect(new).NotTo(BeNil())
	Expect(new.Name).To(Equal(original.Name))
	Expect(new.MapPath).To(Equal(original.MapPath))
	Expect(new.Id).NotTo(BeNil())
}
