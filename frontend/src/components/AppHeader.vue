<template>
  <header class="bg-gray-800 text-white p-4 flex justify-between items-center">
    <h1 class="text-xl font-bold">{{ title }}</h1>
    <div class="flex items-center">
      <span class="mr-4">Welcome, {{ user.username }}</span>
      <button
        @click="logout"
        class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded"
      >
        Logout
      </button>
    </div>
  </header>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

// eslint-disable-next-line
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
});

const router = useRouter();
const user = ref(JSON.parse(localStorage.getItem("user")));

const logout = () => {
  localStorage.removeItem("user");
  router.push("/");
};
</script>
