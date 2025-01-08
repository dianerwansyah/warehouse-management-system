<template>
  <aside class="w-64 bg-gray-800 text-white h-screen p-4">
    <h2 class="text-2xl font-bold mb-4">{{ user.role }} Panel</h2>
    <nav>
      <router-link
        to="/dashboard"
        class="block py-2 px-4 hover:bg-gray-700 rounded"
      >
        Dashboard
      </router-link>

      <router-link
        v-if="allowedRoles.products.includes(user.role)"
        to="/products"
        class="block py-2 px-4 hover:bg-gray-700 rounded"
      >
        Products
      </router-link>
    </nav>
  </aside>
</template>

<script setup>
import { ref } from "vue";

const user = ref(JSON.parse(localStorage.getItem("user")));

const allowedRoles = {
  products: ["admin", "manager"],
};
</script>
