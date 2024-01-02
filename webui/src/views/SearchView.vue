<script>
export default {
	data: function() {
		return {
			errormsg: null,
            searchUsername: "",
            userID: "",
            search: "",
            username:"",
		}
	},
	methods: {

    async searchUser(){
		try {
            this.username=localStorage.getItem("username");
            this.userID=localStorage.getItem("userID");

            if(this.search==""){
                this.errormsg = "search cannot be empty"
            }else{
                console.log(this.search)
                
                this.$router.push({path: "user/" + this.search + "/account"});
            }
        } 

        catch (e) {
				this.errormsg = e.toString();
		}
		
    },
	
    },
}
</script>

<template>
		<div class="ok">
			<h2>Search User</h2>
		
		<div class="gro">
        <input class="text" type="text" v-model="search" placeholder="Search User">
        <button class="btn btn-outline-secondary" id="submit" @click="searchUser" > search  </button>        </div>
        </div>



		
</template>
<style>

.text {
 width: 500px;


}
.search {
 padding: 5px;
 margin-top: 8px;


}
.ok {
  display: flex;
  flex-direction:column;
  align-items: center;
  justify-content: center;
  margin: 10ch;


}
.stuff {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

</style>