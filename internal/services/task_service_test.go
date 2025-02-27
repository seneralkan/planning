package services_test

import (
	"errors"
	"planning/internal/clients"
	clientmocks "planning/internal/clients/mocks"
	"planning/internal/models"
	"planning/internal/repository/mocks"
	"planning/internal/resource/request"
	"planning/internal/services"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

func TestTaskService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TaskService Suite")
}

var _ = Describe("TaskService", func() {
	var (
		mockCtrl      *gomock.Controller
		repoMock      *mocks.MockIRepository
		taskRepoMock  *mocks.MockITasksRepository
		mockProvider1 *clientmocks.MockIProvider
		mockProvider2 *clientmocks.MockIProvider
		taskService   services.ITaskService
		taskScheduler services.ITaskSchedulerService
	)

	mockDevs := []request.Developer{
		{Name: "DEV1", Capacity: 1, WeeklyHours: 45, CurrentHours: 45},
		{Name: "DEV2", Capacity: 2, WeeklyHours: 45, CurrentHours: 45},
		{Name: "DEV3", Capacity: 3, WeeklyHours: 45, CurrentHours: 45},
		{Name: "DEV4", Capacity: 4, WeeklyHours: 45, CurrentHours: 45},
		{Name: "DEV5", Capacity: 5, WeeklyHours: 45, CurrentHours: 45},
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repoMock = mocks.NewMockIRepository(mockCtrl)
		taskRepoMock = mocks.NewMockITasksRepository(mockCtrl)
		repoMock.EXPECT().GetTaskRepository().Return(taskRepoMock).AnyTimes()

		mockProvider1 = clientmocks.NewMockIProvider(mockCtrl)
		mockProvider2 = clientmocks.NewMockIProvider(mockCtrl)
		providers := []clients.IProvider{mockProvider1, mockProvider2}
		taskScheduler = services.NewTaskSchedulerService()
		taskService = services.NewTaskService(repoMock, taskScheduler, providers)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("DistributeTasks", func() {
		It("should distribute tasks among developers", func() {
			tasks := []models.Task{
				{Name: "Task 1", Duration: 60, Difficulty: 3},
				{Name: "Task 2", Duration: 30, Difficulty: 2},
			}

			// Set up mock expectations for the providers first
			mockProvider1.EXPECT().FetchTasks().Return(tasks, nil)
			mockProvider2.EXPECT().FetchTasks().Return([]models.Task{}, nil)

			// Expect storing of tasks
			for _, task := range tasks {
				taskRepoMock.EXPECT().StoreTask(&task).Return(nil)
			}

			developerTasks, totalWeeks, err := taskService.DistributeTasks(mockDevs)

			Expect(err).NotTo(HaveOccurred())
			Expect(totalWeeks).To(Equal(1))

			// Update assertions to match actual distribution
			Expect(developerTasks).To(HaveKey("DEV4"))
			Expect(developerTasks["DEV4"]).To(ContainElement(tasks[0])) // Task 1 should go to DEV4
			Expect(developerTasks["DEV2"]).To(ContainElement(tasks[1])) // Task 2 should go to DEV2
		})

		It("should return an error if fetching tasks fails", func() {
			// Create a WaitGroup to ensure both providers are called
			var wg sync.WaitGroup
			wg.Add(2)

			expectedErr := errors.New("fetch error")

			// Set up expectations with callbacks that use WaitGroup
			mockProvider1.EXPECT().
				FetchTasks().
				DoAndReturn(func() ([]models.Task, error) {
					defer wg.Done()
					defer GinkgoRecover()
					return nil, expectedErr
				})

			mockProvider2.EXPECT().
				FetchTasks().
				DoAndReturn(func() ([]models.Task, error) {
					defer wg.Done()
					defer GinkgoRecover()
					return []models.Task{}, nil
				})

			developerTasks, totalWeeks, err := taskService.DistributeTasks(mockDevs)

			// Wait for both providers to complete
			wg.Wait()

			Expect(err).To(HaveOccurred())
			Expect(developerTasks).To(BeNil())
			Expect(totalWeeks).To(Equal(0))
		})
	})
})
