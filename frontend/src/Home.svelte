<script>
  import { onMount } from "svelte";
  import TreeView from "./TreeView.svelte";

  let godowns = [];
  let items = [];
  let filteredGodowns = [];
  let searchQuery = "";
  let selectedItem = null;
  let loading = true;
  let error = null;

  // Fetch data on mount
  onMount(() => {
    const token = localStorage.getItem("auth_token");
    if (token) {
      fetchData();
    }
  });

  async function fetchData() {
    try {
      const [godownsRes, itemsRes] = await Promise.all([
        fetch("http://localhost:8080/api/godowns"),
        fetch("http://localhost:8080/api/items"),
      ]);

      if (!godownsRes.ok || !itemsRes.ok) {
        throw new Error("Failed to fetch data.");
      }

      godowns = await godownsRes.json();
      items = await itemsRes.json();
      // filteredGodowns = godowns;
      filteredGodowns = getChildren("null"); // Initialize with all godowns
      console.log(filteredGodowns);
    } catch (err) {
      error = err.message;
      console.error("Error fetching data:", err);
    } finally {
      loading = false;
    }
  }

  function filterItems(query) {
    searchQuery = query.toLowerCase();

    return items.filter((item) =>
      item.name.toLowerCase().includes(searchQuery)
    );
  }

  function findTopmostGodown(godown_id) {
    if (godown_id === "null") return godown_id;
    let tmp = godowns.filter((godown) => godown.id === godown_id);
    findTopmostGodown(tmp[0].parent_godown);
  }

  function findTopmostGodowns(filteredItems) {
    let tmp = [];
    for (var i = 0; i < filteredItems.length; i++) {
      tmp.push(findTopmostGodown(filteredItems[i]));
    }
    filteredGodowns = [...new Set(tmp)];
  }

  function getChildren(parentId) {
    return godowns.filter(
      (godown) => godown.parent_godown === parentId
    );
  }

  function getItems(godownId) {
    return items.filter((item) => item.godown_id === godownId);
  }

  function selectItem(item) {
    selectedItem = item;
  }

  function handleSearch(event) {
    const input = event.target;
    if (input && input.value !== undefined) {
      findTopmostGodowns(filterItems(input.value));
    }
  }
</script>

<div class="container">
  <div class="search-bar">
    <input type="text" placeholder="Search items..." on:input={handleSearch} />
  </div>

  <div class="tree-view">
    {#if loading}
      <div class="loading">Loading data...</div>
    {:else if error}
      <div class="error">Error: {error}</div>
    {:else}
      {#each filteredGodowns as godown}
        <TreeView children={godown} {getChildren} {getItems} {selectItem} />
      {/each}
    {/if}
  </div>

  <div class="content">
    {#if selectedItem}
      <div class="content-container">
        <h2>{selectedItem.name}</h2>
        <p><strong>Quantity:</strong> {selectedItem.quantity}</p>
        <p><strong>Category:</strong> {selectedItem.category}</p>
        <p><strong>Price:</strong> ₹{selectedItem.price}</p>
        <p><strong>Status:</strong> {selectedItem.status}</p>
        <p><strong>Brand:</strong> {selectedItem.brand}</p>
        <img src={selectedItem.image_url} alt={selectedItem.name} width="400" />
      </div>
    {/if}
  </div>
</div>

<style>
  .container {
    display: flex;
    height: 100vh;
    overflow: hidden;
  }

  .search-bar {
    width: 400px;
    padding: 1rem;
    background-color: #f1f1f1;
    border-bottom: 1px solid #ddd;
  }

  .search-bar input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  .tree-view {
    width: 400px;
    position: fixed;
    left: 0;
    top: 80px; /* Adjust to fit search bar */
    bottom: 0;
    background-color: #f1f1f1;
    border-right: 1px solid #ddd;
    overflow-y: auto;
    padding: 1rem;
  }

  .content {
    flex: 1;
    margin-left: 400px;
    padding: 2rem;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .content-container {
    max-width: 600px;
    width: 100%;
    background-color: #f9f9f9;
    border-radius: 12px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 2rem;
  }

  .content img {
    display: block;
    margin: 1rem auto;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .loading,
  .error {
    text-align: center;
    margin-top: 2rem;
    font-size: 1.5rem;
  }
</style>
