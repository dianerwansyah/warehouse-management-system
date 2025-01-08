import { createRouter, createWebHistory } from "vue-router";
import LoginView from "@/views/LoginView.vue";
import DashboardView from "@/views/DashboardView.vue";
import ProductsView from "@/views/ProductsView.vue";

const allowedRoles = {
  products: ["admin", "manager"],
};

const getUserFromLocalStorage = () => {
  return JSON.parse(localStorage.getItem("user"));
};

const routes = [
  {
    path: "/",
    name: "Login",
    component: LoginView,
    beforeEnter: (to, from, next) => {
      const user = getUserFromLocalStorage();
      if (user) {
        next({ name: "Dashboard" });
      } else {
        next();
      }
    },
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: DashboardView,
  },
  {
    path: "/products",
    name: "Products",
    component: ProductsView,
    meta: { allowedRoles: allowedRoles.products },
  },
];

// eslint-disable-next-line
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const user = getUserFromLocalStorage();
  if (!user && to.name !== "Login") {
    next({ name: "Login" });
  } else {
    const { allowedRoles } = to.meta;
    if (allowedRoles && !allowedRoles.includes(user?.role)) {
      next({ name: "Dashboard" });
    } else {
      next();
    }
  }
});

export default router;
