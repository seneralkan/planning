package repository_test

import (
	"database/sql"
	"os"
	"testing"

	"planning/internal/models"
	"planning/internal/repository"
	"planning/pkg/sqlite"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = Describe("Repository", func() {
	var (
		db             *sql.DB
		sqliteInstance sqlite.ISqliteInstance
		repo           repository.IRepository
		taskRepo       repository.ITasksRepository
	)

	BeforeEach(func() {
		var err error
		sqliteInstance, err = sqlite.NewSqliteInstance("test")
		Expect(err).NotTo(HaveOccurred())

		db = sqliteInstance.Database()
		taskRepo = repository.NewTaskRepository(sqliteInstance)
		repo = repository.NewRepository(taskRepo)

		_, err = db.Exec(`DROP TABLE IF EXISTS tasks`)
		Expect(err).NotTo(HaveOccurred())

		_, err = db.Exec(`
			CREATE TABLE tasks (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				duration INTEGER,
				difficulty INTEGER
			);
		`)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		sqliteInstance.Close()
		os.Remove("test.db")
	})

	Describe("Repository", func() {
		It("should return the task repository", func() {
			actual := repo.GetTaskRepository()
			Expect(actual).To(Equal(taskRepo))
		})
	})

	Describe("StoreTask", func() {
		It("should store a task in the database", func() {
			task := &models.Task{Name: "Test Task", Duration: 60, Difficulty: 3}
			err := repo.GetTaskRepository().StoreTask(task)
			Expect(err).NotTo(HaveOccurred())

			var count int
			err = db.QueryRow(`SELECT COUNT(*) FROM tasks WHERE name = ?`, task.Name).Scan(&count)
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(1))
		})
	})

	Describe("GetTasks", func() {
		It("should retrieve tasks from the database", func() {
			_, err := db.Exec(`INSERT INTO tasks (name, duration, difficulty) VALUES (?, ?, ?)`, "Test Task", 60, 3)
			Expect(err).NotTo(HaveOccurred())

			tasks, err := repo.GetTaskRepository().GetTasks()
			Expect(err).NotTo(HaveOccurred())
			Expect(tasks).To(HaveLen(1))
			Expect(tasks[0].Name).To(Equal("Test Task"))
		})
	})
})
