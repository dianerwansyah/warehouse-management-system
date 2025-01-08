<template>
  <div
    v-if="showEditProductModal"
    class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
  >
    <div class="bg-white p-6 rounded shadow-md w-96">
      <h2 class="text-xl font-bold mb-4">Edit Product</h2>
      <input
        v-model="Product.name"
        placeholder="Product Name"
        class="border p-2 w-full mb-2"
      />
      <input
        v-model="Product.sku"
        placeholder="SKU"
        class="border p-2 w-full mb-2"
      />
      <input
        v-model="Product.qty"
        type="number"
        placeholder="Quantity"
        class="border p-2 w-full mb-2"
      />
      <input
        v-model="Product.location"
        placeholder="Location"
        class="border p-2 w-full mb-2"
      />
      <select v-model="Product.status" class="border p-2 w-full mb-4">
        <option value="available">Available</option>
        <option value="out of stock">Out of Stock</option>
      </select>
      <button
        @click="handleEditProduct"
        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
      >
        Save Changes
      </button>
      <button
        @click="closeModal"
        class="bg-gray-300 hover:bg-gray-400 text-black font-bold py-2 px-4 rounded ml-2"
      >
        Cancel
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";

// eslint-disable-next-line
const props = defineProps({
  showEditProductModal: Boolean,
  productToEdit: Object,
});

// eslint-disable-next-line
const emit = defineEmits(["close", "refresh"]);

const Product = ref({ ...props.productToEdit });

const handleEditProduct = async () => {
  if (
    !Product.value.name ||
    !Product.value.sku ||
    Product.value.qty <= 0 ||
    !Product.value.location ||
    !Product.value.status
  ) {
    alert("All fields are required and quantity must be greater than zero.");
    return;
  }

  try {
    await axios.put("/products/update", Product.value);

    emit("refresh");
    emit("close");
  } catch (error) {
    console.error("Error editing product:", error);
  }
};

const closeModal = () => {
  emit("close");
};
</script>
