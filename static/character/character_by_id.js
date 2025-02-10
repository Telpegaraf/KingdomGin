function showMasteryOptions(skillId, currentMastery) {
    const masteryLevels = ["None", "Trained", "Expert", "Master", "Legendary"];
    let options = masteryLevels.map(level =>
        `<option value="${level}" ${level === currentMastery ? 'selected' : ''}>${level}</option>`
    ).join('');

    const selectHtml = `<select id="select-${skillId}" onchange="updateMastery(${skillId}, this.value)">
                                    ${options}
                                </select>`;

    const element = document.querySelector(`.skill-mastery[data-id="${skillId}"]`);
    if (element) {
        element.innerHTML = selectHtml;
        const selectElement = document.getElementById(`select-${skillId}`);
        if (selectElement) {
            selectElement.focus();
            selectElement.style.display = 'inline-block';

            // Добавляем обработчик события для скрытия меню при уходе курсора
            selectElement.addEventListener('mouseleave', () => {
                selectElement.style.display = 'none';
                element.innerHTML = currentMastery; // Возвращаем текст
            });
        }
    } else {
        console.error(`Element with skillId ${skillId} not found.`);
    }
}

function updateMastery(skillId, newMastery) {
    fetch(`/character-skill/${skillId}`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ mastery: newMastery })
    })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                document.querySelector(`.skill-mastery[data-id="${skillId}"]`).innerText = newMastery;
            }
        })
        .catch(error => console.error('Error:', error));
}