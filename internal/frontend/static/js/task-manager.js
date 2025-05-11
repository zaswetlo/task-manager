const taskManager = {
    init() {
        this.fetchTasks();
        this.updateProgress();
    },

    async fetchTasks() {
        try {
            const response = await fetch('/api/tasks');
            const tasks = await response.json();
            const taskList = document.getElementById('taskList');
            taskList.innerHTML = '';
            
            tasks.forEach(task => {
                const li = this.createTaskElement(task);
                taskList.appendChild(li);
            });
        } catch (error) {
            console.error('Error fetching tasks:', error);
        }
    },

    async updateProgress() {
        try {
            const response = await fetch('/api/tasks/progress');
            const data = await response.json();
            const progress = Math.round(data.progress);
            
            document.getElementById('progressFill').style.width = `${progress}%`;
            document.getElementById('progressText').textContent = `${progress}% Complete`;
        } catch (error) {
            console.error('Error updating progress:', error);
        }
    },

    async addTask() {
        const input = document.getElementById('taskInput');
        const title = input.value.trim();
        
        if (!title) return;

        try {
            const response = await fetch('/api/tasks', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ title }),
            });

            if (response.ok) {
                input.value = '';
                this.fetchTasks();
                this.updateProgress();
            }
        } catch (error) {
            console.error('Error adding task:', error);
        }
    },

    createTaskElement(task) {
        const li = document.createElement('li');
        li.className = 'task-item';
        if (task.done) li.classList.add('completed');

        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.checked = task.done;
        checkbox.onchange = () => this.toggleTask(task.id, checkbox.checked);

        const content = document.createElement('div');
        content.className = 'task-content';

        const title = document.createElement('div');
        title.className = 'task-title';
        title.textContent = task.title;

        const description = document.createElement('div');
        description.className = 'task-description';
        description.textContent = task.description || 'No description';

        const actions = document.createElement('div');
        actions.className = 'task-actions';

        const editButton = document.createElement('button');
        editButton.textContent = 'Edit';
        editButton.onclick = () => this.toggleEditForm(task.id);

        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.className = 'delete';
        deleteButton.onclick = () => this.deleteTask(task.id);

        const editForm = document.createElement('div');
        editForm.className = 'edit-form';
        editForm.id = `edit-form-${task.id}`;
        editForm.innerHTML = `
            <input type="text" id="edit-title-${task.id}" value="${task.title}" placeholder="Task title">
            <textarea id="edit-description-${task.id}" placeholder="Task description">${task.description || ''}</textarea>
            <button onclick="taskManager.updateTaskDetails(${task.id})">Save</button>
            <button onclick="taskManager.toggleEditForm(${task.id})">Cancel</button>
        `;

        actions.appendChild(editButton);
        actions.appendChild(deleteButton);
        content.appendChild(title);
        content.appendChild(description);
        content.appendChild(editForm);
        li.appendChild(checkbox);
        li.appendChild(content);
        li.appendChild(actions);

        return li;
    },

    toggleEditForm(taskId) {
        const form = document.getElementById(`edit-form-${taskId}`);
        form.classList.toggle('active');
    },

    async updateTaskDetails(taskId) {
        const title = document.getElementById(`edit-title-${taskId}`).value.trim();
        const description = document.getElementById(`edit-description-${taskId}`).value.trim();

        if (!title) return;

        try {
            const response = await fetch(`/api/tasks/${taskId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ title, description }),
            });

            if (response.ok) {
                this.fetchTasks();
            }
        } catch (error) {
            console.error('Error updating task:', error);
        }
    },

    async toggleTask(id, done) {
        try {
            const response = await fetch(`/api/tasks/${id}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ done }),
            });

            if (response.ok) {
                this.fetchTasks();
                this.updateProgress();
            }
        } catch (error) {
            console.error('Error updating task:', error);
        }
    },

    async deleteTask(id) {
        if (!confirm('Are you sure you want to delete this task?')) return;

        try {
            const response = await fetch(`/api/tasks/${id}`, {
                method: 'DELETE',
            });

            if (response.ok) {
                this.fetchTasks();
                this.updateProgress();
            }
        } catch (error) {
            console.error('Error deleting task:', error);
        }
    }
};

// Initialize the task manager when the page loads
window.onload = () => taskManager.init(); 