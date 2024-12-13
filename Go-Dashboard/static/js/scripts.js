// Define the base API endpoint for server communication
const API_BASE = "/api"; // Replace with your server endpoint if different

// Fetch data from the JSON file via the server
async function fetchData() {
  const response = await fetch(`${API_BASE}/data`); // Make a GET request to fetch data
  const data = await response.json(); // Parse the response as JSON
  categories = data.categories || []; // Assign categories from the response, default to an empty array
  bookmarks = data.bookmarks || []; // Assign bookmarks from the response, default to an empty array
  renderCategories(); // Update the UI with fetched categories
  renderBookmarks(); // Update the UI with fetched bookmarks
}

// Save the current state of categories and bookmarks to the server
async function saveData() {
  const response = await fetch(`${API_BASE}/data`, {
    method: "POST", // Use POST to update data on the server
    headers: { "Content-Type": "application/json" }, // Inform the server we're sending JSON
    body: JSON.stringify({ categories, bookmarks }), // Send categories and bookmarks as JSON
  });
  if (response.ok) {
    location.reload(); // Refresh the page to reflect changes
  } else {
    alert("Failed to save data. Please try again."); // Show an error message if saving fails
  }
}

// Render the list of categories in the UI
function renderCategories() {
  const categoriesContainer = document.getElementById("categories-container"); // Get the categories container
  const categorySelect = document.getElementById("bookmark-category"); // Get the category dropdown

  categoriesContainer.innerHTML = ""; // Clear existing categories
  categorySelect.innerHTML = ""; // Clear existing dropdown options

  // Loop through categories and render each one
  categories.forEach((category, index) => {
    const categoryDiv = document.createElement("div"); // Create a div for each category
    categoryDiv.innerHTML = `
      ${category} 
      <button onclick="removeCategory(${index})">Remove</button> 
    `; // Add the category name and a remove button
    categoriesContainer.appendChild(categoryDiv); // Add the category div to the container

    const option = document.createElement("option"); // Create an option for the dropdown
    option.value = category; // Set the value to the category name
    option.textContent = category; // Set the display text
    categorySelect.appendChild(option); // Add the option to the dropdown
  });
}

// Add a new category
async function addCategory() {
  const categoryName = document.getElementById("category-name").value.trim(); // Get the input value and trim whitespace
  if (!categoryName) {
    alert("Category name is required!"); // Validate input
    return;
  }
  if (categories.includes(categoryName)) {
    alert("Category already exists!"); // Prevent duplicate categories
    return;
  }

  categories.push(categoryName); // Add the new category to the list
  await saveData(); // Save the updated categories to the server
}

// Remove a category if it has no associated bookmarks
async function removeCategory(index) {
  const categoryToRemove = categories[index]; // Get the category to remove
  const hasBookmarks = bookmarks.some(bookmark => bookmark.category === categoryToRemove); // Check if the category has associated bookmarks

  if (hasBookmarks) {
    alert("Cannot remove category with associated bookmarks!"); // Show an error if bookmarks exist
    return;
  }

  categories.splice(index, 1); // Remove the category from the list
  await saveData(); // Save the updated categories to the server
}

// Render the list of bookmarks grouped by category
function renderBookmarks() {
  const container = document.getElementById("bookmarks-container"); // Get the bookmarks container
  container.innerHTML = ""; // Clear existing bookmarks

  const groupedBookmarks = bookmarks.reduce((groups, bookmark) => {
    if (!groups[bookmark.category]) groups[bookmark.category] = []; // Initialize the category group if it doesn't exist
    groups[bookmark.category].push(bookmark); // Add the bookmark to the appropriate group
    return groups;
  }, {}); // Group bookmarks by category

  // Loop through categories and render bookmarks for each
  for (const category in groupedBookmarks) {
    const categoryDiv = document.createElement("div"); // Create a div for each category
    categoryDiv.innerHTML = `<h3>${category}</h3>`; // Add the category header
    groupedBookmarks[category].forEach((bookmark, index) => {
      const bookmarkEl = document.createElement("div"); // Create a div for each bookmark
      bookmarkEl.innerHTML = `
        <a href="${bookmark.url}" target="_blank">${bookmark.title}</a>
        <button onclick="editBookmark(${index})">Edit</button>
        <button onclick="removeBookmark(${index})">Remove</button>
      `; // Add the bookmark title, edit, and remove buttons
      categoryDiv.appendChild(bookmarkEl); // Add the bookmark div to the category div
    });
    container.appendChild(categoryDiv); // Add the category div to the container
  }
}

// Add a new bookmark
async function addBookmark() {
  const title = document.getElementById("bookmark-title").value.trim(); // Get the bookmark title
  const url = document.getElementById("bookmark-url").value.trim(); // Get the bookmark URL
  const category = document.getElementById("bookmark-category").value; // Get the selected category

  if (!title || !url) {
    alert("Title and URL are required!"); // Validate input
    return;
  }

  bookmarks.push({ title, url, category }); // Add the new bookmark to the list
  await saveData(); // Save the updated bookmarks to the server
}

// Remove a bookmark
async function removeBookmark(index) {
  bookmarks.splice(index, 1); // Remove the bookmark from the list
  await saveData(); // Save the updated bookmarks to the server
}

// Edit an existing bookmark
async function editBookmark(index) {
  const bookmark = bookmarks[index]; // Get the bookmark to edit
  const newTitle = prompt("Edit Title", bookmark.title); // Prompt for a new title
  const newUrl = prompt("Edit URL", bookmark.url); // Prompt for a new URL
  const newCategory = prompt("Edit Category", bookmark.category); // Prompt for a new category

  if (newTitle && newUrl && newCategory) {
    bookmarks[index] = { title: newTitle, url: newUrl, category: newCategory }; // Update the bookmark
    await saveData(); // Save the updated bookmarks to the server
  }
}

// Handle menu selection to show the appropriate form
function handleMenuChange() {
  const menuSelect = document.getElementById("menu-select").value; // Get the selected menu option
  const categorySection = document.getElementById("add-category-section"); // Get the add-category section
  const bookmarkSection = document.getElementById("add-bookmark-section"); // Get the add-bookmark section

  // Hide all sections initially
  categorySection.classList.add("hidden");
  bookmarkSection.classList.add("hidden");

  // Show the selected section
  if (menuSelect === "add-category") {
    categorySection.classList.remove("hidden");
  } else if (menuSelect === "add-bookmark") {
    bookmarkSection.classList.remove("hidden");
  }
}

// Fetch data from the server and render the UI when the page loads
fetchData();
