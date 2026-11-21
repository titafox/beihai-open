import Vue from 'vue'
import Vuex from 'vuex'


Vue.use(Vuex)


const storea = new Vuex.Store({
	state: {
		
		userInfo:{},
		

	},
	mutations: { //添加mutations
		
		updataUser(state,data){
		  state.userInfo=data; 	
		}
		
		
		

	}
})

// _self.$store.commit('updateuser', obja);
// _self.$store.commit('updateact', objb);

export default storea
