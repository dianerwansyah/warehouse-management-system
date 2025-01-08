<template>
  <div class="flex">
    <NavSidebar />
    <div class="flex-1 p-4">
      <AppHeader title="Dashboard" />
      <main class="mt-4">
        <div class="mb-4">
          <h2 class="text-2xl font-semibold mb-2">Product Overview</h2>
          <div class="grid grid-cols-3 gap-4">
            <div class="bg-white p-4 rounded shadow">
              <h3 class="font-bold">Total Products</h3>
              <p class="text-3xl">{{ totalProducts }}</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
              <h3 class="font-bold">Total Stock</h3>
              <p class="text-3xl">{{ totalStock }}</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
              <h3 class="font-bold">Low Stock Items</h3>
              <p class="text-3xl">{{ lowStockCount }}</p>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import NavSidebar from "../components/NavSidebar.vue";
import AppHeader from "../components/AppHeader.vue";
import axios from "axios";

const totalProducts = ref(0);
const totalStock = ref(0);
const lowStockCount = ref(0);

const fetchProductOverview = async () => {
  try {
    const response = await axios.get("/products/gets");
    const products = response.data;
    totalProducts.value = products.length;
    totalStock.value = products.reduce((sum, product) => sum + product.qty, 0);
    lowStockCount.value = products.filter((product) => product.qty < 5).length;
  } catch (error) {
    console.error("Error fetching products:", error);
  }
};

onMounted(() => {
  fetchProductOverview();
});
</script>
