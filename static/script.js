// Show detail view when an image is clicked
function reformat(data) {
    const subboxes = [];
    data.forEach((element, idx) => {
        const folder_name = element.folder ?? 'root'
        const child = subboxes.find(sb => sb.name === folder_name);
        const img_obj = {
            name: element.name,
            id: idx,
        };
        if (child) {
            child.images.push(img_obj)
        } else {
            subboxes.push({
                name: folder_name,
                images: [img_obj],
            });
        }
    });
    return subboxes;
}

document.addEventListener('DOMContentLoaded', function () {
    const mainBox = document.getElementById('main-box');
    const detailView = document.getElementById('detail-view');

    // Fetch image list from /imagelist
    fetch('/imagelist')
        .then(response => response.json())
        .then(data => {
            data = reformat(data);
            // Create main box and sub-boxes
            data.forEach(subBoxData => {
                const subBox = document.createElement('div');
                subBox.className = 'sub-box';
                subBox.innerHTML = `<div>${subBoxData.name}</div>`;
                
                // Create images inside sub-box
                subBoxData.images.forEach(imgObj => {
                    console.log(imgObj);
                    const image = document.createElement('img');
                    image.src = `/image/${imgObj.id}`;
                    image.className = 'image';
                    image.addEventListener('click', () => showDetailView(image.src));
                    subBox.appendChild(image);
                });

                mainBox.appendChild(subBox);
            });
        });

    // Show detail view when an image is clicked
    function showDetailView(imageUrl) {
        const image = document.createElement('img');
        image.src = imageUrl;
        detailView.innerHTML = '';
        detailView.appendChild(image);
        detailView.style.display = 'flex';
        detailView.addEventListener('click', hideDetailView);
    }

    // Hide detail view when clicked outside the image
    function hideDetailView() {
        detailView.style.display = 'none';
        detailView.removeEventListener('click', hideDetailView);
    }

    // Update layout when the window is resized
    window.addEventListener('resize', updateLayout);

    function updateLayout() {
        // Reset the main box's display property
        mainBox.style.display = 'none';
        mainBox.offsetHeight; // Trigger a reflow
        mainBox.style.display = 'flex';
    }
});
