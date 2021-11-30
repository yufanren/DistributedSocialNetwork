import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "../views/Home"
import Register from "../views/user/Register"
import Login from "../views/user/Login"
import BlogDetail from "../views/blog/Detail"
import BlogEdit from "../views/blog/Edit"
import UserProfile from "../views/user/Profile"

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/blog/create',
    name: 'BlogCreate',
    component: BlogEdit
  },
  {
    path: '/blog/:id',
    name: 'BlogDetail',
    component: BlogDetail
  },
  {
    path: '/user/:username',
    name: 'UserProfile',
    component: UserProfile
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
