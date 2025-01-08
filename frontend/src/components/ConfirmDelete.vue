<template>
  <div
    v-if="isVisible"
    class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
  >
    <div class="bg-white p-6 rounded shadow-md w-96">
      <h2 class="text-xl font-bold mb-4">Confirm Deletion</h2>
      <p class="mb-4">Are you sure you want to delete this item?</p>
      <div class="flex justify-center gap-4">
        <button
          @click="hide"
          class="bg-gray-300 hover:bg-gray-400 text-black font-bold py-2 px-4 rounded"
        >
          Cancel
        </button>
        <button
          @click="confirmDelete"
          class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded"
        >
          Confirm
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, defineExpose, defineEmits } from "vue";

const emit = defineEmits(["confirm", "cancel"]);
const isVisible = ref(false);

function show() {
  isVisible.value = true;
}

function hide() {
  // eslint-disable-next-line
  emit("cancel");
  isVisible.value = false;
}

function confirmDelete() {
  // eslint-disable-next-line
  emit("confirm");
  hide();
}

// eslint-disable-next-line
defineExpose({
  show,
});
</script>
