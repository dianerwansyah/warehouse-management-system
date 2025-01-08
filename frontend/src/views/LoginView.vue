<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-md w-96">
      <h2 class="text-2xl font-bold text-center mb-6">Login</h2>
      <form @submit.prevent="login">
        <div class="mb-4">
          <input
            v-model="username"
            type="text"
            placeholder="Username"
            required
            class="w-full p-2 border border-gray-300 rounded"
          />
        </div>
        <div class="mb-4">
          <input
            v-model="password"
            type="password"
            placeholder="Password"
            required
            class="w-full p-2 border border-gray-300 rounded"
          />
        </div>
        <button
          type="submit"
          class="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
        >
          Login
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

const username = ref("");
const password = ref("");
const router = useRouter();

const login = async () => {
  try {
    const response = await axios.post("/login", {
      username: username.value,
      password: password.value,
    });
    const userData = {
      token: response.data.token,
      username: response.data.username,
      role: response.data.role,
    };

    localStorage.setItem("user", JSON.stringify(userData));

    router.push({ name: "Dashboard" });
  } catch (error) {
    console.error(
      "Login error:",
      error.response ? error.response.data : error.message
    );
    alert("Invalid Login ID or Password");
  }
};
</script>

<style scoped></style>
