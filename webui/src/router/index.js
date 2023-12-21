import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import login from '../views/LoginView.vue'
import AccountView from '../views/AccountView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: login},
		{path: '/session', component: HomeView},
		{path: '/account', component: AccountView},
	//	{path: '/users/:Username/', component: AccountView},

	]
})

export default router
