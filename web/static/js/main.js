let developers = [];

$(document).ready(function() {
    const form = $('#developerForm');
    
    // Initialize developers list
    updateDevelopersList();
    
    form.on('submit', function(event) {
        event.preventDefault();
        
        const name = $('#developerName').val().trim();
        const weeklyHours = parseInt($('#weeklyHours').val());
        const capacity = parseInt($('#capacity').val());
        
        if (!name || !weeklyHours || !capacity) {
            alert('Please fill all fields');
            return;
        }
        
        if (weeklyHours < 1) {
            alert('Weekly hours must be greater than 0');
            return;
        }
        
        if (capacity < 1 || capacity > 10) {
            alert('Capacity must be between 1 and 10');
            return;
        }
        
        const developer = {
            name: name,
            weeklyHours: weeklyHours,
            currentHours: weeklyHours,
            capacity: capacity
        };
        
        // Add developer and update UI
        developers.push(developer);
        updateDevelopersList();
        
        // Reset form
        form[0].reset();
        console.log('Developer added:', developer);
        console.log('Current developers:', developers);
    });

    // Attach click handler to the developers list container for remove buttons
    $('#developersList').on('click', '.remove-developer', function() {
        const index = $(this).data('index');
        removeDeveloper(index);
    });

    $('#distributeTasks').click(function() {
        if (developers.length === 0) {
            alert('Please add at least one developer');
            return;
        }

        const requestBody = {
            developers: developers.map(dev => ({
                name: dev.name,
                weeklyHours: dev.weeklyHours,
                currentHours: dev.weeklyHours, // Initially same as weeklyHours
                capacity: dev.capacity
            }))
        };

        $.ajax({
            url: '/api/tasks/schedule',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(requestBody),
            success: function(response) {
                displayResults(response);
            },
            error: function(error) {
                alert('Error distributing tasks: ' + error.responseJSON.error);
            }
        });
    });
});

function updateDevelopersList() {
    const list = $('#developersList');
    list.empty();
    
    if (developers.length === 0) {
        list.append('<li class="list-group-item text-muted">No developers added yet</li>');
        return;
    }
    
    developers.forEach((dev, index) => {
        list.append(`
            <li class="list-group-item d-flex justify-content-between align-items-center">
                <div>
                    <strong>${dev.name}</strong> - ${dev.weeklyHours} hours/week
                    <span class="badge bg-primary ms-2">Capacity: ${dev.capacity}</span>
                </div>
                <button class="btn btn-danger btn-sm remove-developer" data-index="${index}">Remove</button>
            </li>
        `);
    });
}

function removeDeveloper(index) {
    if (index >= 0 && index < developers.length) {
        developers.splice(index, 1);
        updateDevelopersList();
    }
}

function displayResults(response) {
    const results = $('#results');
    results.empty();
    
    results.append(`<h4>Total Weeks: ${response.totalWeeks}</h4>`);
    
    Object.entries(response.schedule).forEach(([developer, tasks]) => {
        const tasksList = tasks.map(task => 
            `<li class="list-group-item">
                ${task.name} (Difficulty: ${task.difficulty}, Duration: ${task.duration} hours)
             </li>`
        ).join('');
        
        results.append(`
            <div class="card mb-3">
                <div class="card-header">${developer}</div>
                <ul class="list-group list-group-flush">
                    ${tasksList}
                </ul>
            </div>
        `);
    });
}
