func aiSuggestTask(task string) string {
	// Simple AI-mock: Append "AI-enhanced"
	return task + " (AI-enhanced)"
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock AI-enhanced task
	task.Title = aiSuggestTask(task.Title)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := taskCollection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating task"})
		return
	}

	broadcast <- task
	c.JSON(http.StatusOK, task)
}
