import {createRouter, createWebHashHistory} from 'vue-router'
import stream from '../views/StreamView.vue'
import login from '../views/LogView.vue'
import profile from '../views/MyProfileView.vue'
import userAccount from '../views/UserProfileView.vue'
import search from '../views/SearchView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: login},
		{path: '/session', component: stream},
		{path: '/profile', component: profile},
		{path: '/user/:username/account', component: userAccount},
		{path: '/home', component: stream},
		{path: '/search', component: search},
		
	]
})

export default router
