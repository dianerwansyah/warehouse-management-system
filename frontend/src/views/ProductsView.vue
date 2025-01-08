<template>
  <div class="flex">
    <NavSidebar />
    <div class="flex-1 p-4">
      <AppHeader title="Products" />
      <div class="mt-4 mb-4 flex items-center">
        <button
          @click="showAddProductModal = true"
          class="ml-auto bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        >
          Add Product
        </button>
      </div>

      <table class="min-w-full bg-white border border-gray-300">
        <thead>
          <tr>
            <th class="border px-4 py-2">No</th>
            <th class="border px-4 py-2">Name</th>
            <th class="border px-4 py-2">SKU</th>
            <th class="border px-4 py-2">Quantity</th>
            <th class="border px-4 py-2">Location</th>
            <th class="border px-4 py-2">Status</th>
            <th class="border px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(product, index) in products" :key="product.id">
            <td class="border px-4 py-2">{{ index + 1 }}</td>
            <td class="border px-4 py-2">{{ product.name }}</td>
            <td class="border px-4 py-2">{{ product.sku }}</td>
            <td class="border px-4 py-2">{{ product.qty }}</td>
            <td class="border px-4 py-2">{{ product.location }}</td>
            <td class="border px-4 py-2">{{ product.status }}</td>
            <td class="border px-4 py-2">
              <button
                @click="openEdit(product)"
                class="bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-1 px-2 rounded mr-2"
              >
                Edit
              </button>
              <button
                @click="openConfirmDelete(product.id)"
                class="bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded"
              >
                Delete
              </button>
              <button
                @click="downloadBarcode(product.sku)"
                class="bg-gray-500 hover:bg-gray-600 text-white font-bold py-1 px-2 rounded ml-2"
              >
                Download Barcode
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <ProductAdd
        v-if="showAddProductModal"
        :newProduct="newProduct"
        :showAddProductModal="showAddProductModal"
        @close="showAddProductModal = false"
        @refresh="fetchProducts"
      />

      <ProductEdit
        v-if="showEditProductModal"
        :showEditProductModal="showEditProductModal"
        :productToEdit="productToEdit"
        @close="showEditProductModal = false"
        @refresh="fetchProducts"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import NavSidebar from "../components/NavSidebar.vue";
import AppHeader from "../components/AppHeader.vue";
import ProductAdd from "./widget/ProductAdd.vue";
import ProductEdit from "./widget/ProductEdit.vue";
import JsBarcode from "jsbarcode"; // Import JsBarcode

const products = ref([]);
const showAddProductModal = ref(false);
const showEditProductModal = ref(false);
const newProduct = ref({
  name: "",
  sku: "",
  qty: 0,
  location: "",
  status: "available",
});
const productToEdit = ref(null);

const fetchProducts = async () => {
  try {
    const response = await axios.get("/products/gets");
    products.value = response.data;
  } catch (error) {
    console.error("Error fetching products:", error);
  }
};

const openConfirmDelete = (id) => {
  const confirmDelete = confirm(
    "Are you sure you want to delete this product?"
  );
  if (confirmDelete) {
    deleteProduct(id);
  }
};

const deleteProduct = async (id) => {
  try {
    await axios.delete(`/products/delete/${id}`);
    fetchProducts();
  } catch (error) {
    console.error("Error deleting product:", error);
  }
};

const openEdit = (product) => {
  showEditProductModal.value = true;
  productToEdit.value = { ...product };
};

const downloadBarcode = (sku) => {
  const canvas = document.createElement("canvas");

  JsBarcode(canvas, sku, {
    format: "CODE128",
    width: 2,
    height: 100,
    displayValue: true,
  });

  const imgData = canvas.toDataURL("image/png");
  const link = document.createElement("a");
  link.setAttribute("href", imgData);
  link.setAttribute("download", `${sku}_barcode.png`);
  link.click();
};

onMounted(fetchProducts);
</script>
