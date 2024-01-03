<script>
import LoadingSpinner from "../components/LoadingSpinner.vue";
export default {
	data: function() {
		return {
            errormsg: null,
           
            userUsername: localStorage.getItem("username"),
            newUser: "",
            userID: localStorage.getItem("userID"),
            followerCount: null,
            followingCount: null,
            followers:[],
            following:[],
            bans:[],
            bannedCount:null,
            photoCount:null,
            isFollow: false,
            isBan: false,
            isLoading: true,
            likeID: 0,

        
            photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLiked:true,
                comment: "",
                comments:[{
                        userID:0,
                        username: "",
                        comment:"",

                    }],
                }
            ],
            }

        
		},
     created(){

        this.AccountInfo();
        //this.getPhotos();
        
    },

	methods: {

      

        mounted(){
        this.AccountInfo()
       
    },

    async popupComment() {
        var modal = document.getElementById("commentModal");
   
        modal.style.display = "block";


    },
    async closeModal() {
        var modal = document.getElementById("commentModal");
        var followmodal = document.getElementById("followerModal");
        var followingmodal = document.getElementById("followingModal");
        var banModal = document.getElementById("banModal");

        followmodal.style.display = "none";
        followingmodal.style.display = "none";
        banModal.style.display = "none";
        modal.style.display = "none";
    },

		 async AccountInfo() {
            
			try {
                console.log("idk");
                this.isLoading=true;
                this.userUsername=localStorage.getItem("username");
                this.userID=localStorage.getItem("userID");
                console.log(this.userUsername);
                let response=await this.$axios.get("/username/"+ this.userUsername + "/profile")
                    this.followerCount=response.data.followed_count;
                    this.followingCount=response.data.following_count;
                    this.photoCount=response.data.pic_numb;
                    this.bannedCount=response.data.ban_count;
                    this.bans=response.data.bans;

                console.log("bannn: " + this.bans);
                
               

                if(this.photoCount>0){
                    this.getPhotos();
                };
                
                console.log("plsplspls work");

            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async refresh() {
		
			this.AccountInfo();
            
	},


    async LikePhoto(photoID){

console.log("putting like")

        let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/like");
    console.log(response.data);
    this.getPhotos();
    
    

},
async unLikePhoto(photoID){

console.log("deleting like")
let response=await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/like");
console.log(response.data);

this.getPhotos();

},


    async updateUsername(){

        if(this.newUser==""){
            this.errormsg="you must enter a username!!!"

        }else{
            console.log(this.newUser);
            console.log(this.userID);
            let response=await this.$axios.put("/user/"+ this.userID + "/setusername/" + this.newUser);
            localStorage.setItem("username", response.data);
            console.log(response);
            this.refresh();
        }


    },
	
    async getFollower(){
        try{     
           
            let response=await this.$axios.get("/user/"+ this.userID + "/followers");
            console.log("followers: " +response.data );
            this.followers=response.data;
            console.log(this.followers);
            var modal = document.getElementById("followerModal");
            modal.style.display = "block"; 
        }
       
    
     catch(e){
            this.errormsg=e.toString();
        }


    },
    async getFollowing(){
        try{     
           
            let response=await this.$axios.get("/user/"+ this.userID + "/followings");
            console.log("followings: " +response.data );
            this.following=response.data;
            console.log(this.following);
            var modal = document.getElementById("followingModal");
            modal.style.display = "block";
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async getBans(){
        try{    
            var modal = document.getElementById("banModal");
            modal.style.display = "block";
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },

    async uploadPhoto(){
        try{     
            console.log("hahaha")
                const uploadpic=this.$refs.pico.files[0];
                // read.onload= async () =>{};
                let response= await this.$axios.post("/photo/upload/" + this.userID, uploadpic);
                this.AccountInfo();
                this.refresh();
            }
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async deletePhoto(photoID) {
        let response= await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/remove");
        this.getPhotos();
       
    },

       
    async getPhotos(){
           
                let response=await this.$axios.get("/user/" + this.userID + "/photo/" + this.userID );
                console.log(response.data);
                this.photos=response.data;
                console.log(this.photos);
                console.log(this.photos[0].photo)
                console.log(this.photos.length);
                for (let i=0;i<this.photos.length;i++){
                    
                    this.photos[i].photo= 'data:image/png;base64,'+ this.photos[i].photo;
                    console.log("omg   "+this.photos[i].photo);
                    console.log("slay   "+this.photos[i].photoID);
                    console.log("slay   "+this.photos[i].likeCount);
                      
                }
                
        
    },

    async uploadComment(photoID,yuhcomment){
        console.log(yuhcomment);
           
           let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/comment", { comment: yuhcomment });
           console.log(response.data);
           this.AccountInfo();
           this.getPhotos();
},




    
 

},
}

 

</script>
<template>
    <head><meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></head>
  
    

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

  
    <div class="upload">
    <div class="btn-group btn-group-sm">
					<input type="file" accept="image/*" class="btn btn-xs" id="inputImage" ref="pico" @change="uploadFile" >

                    <button class="btn btn-default " @click="uploadPhoto">Upload</button>
				</div>
    </div>
   
    <div class="umidk">
    <div class="btn-group">
 
            <button class="btn" @click="getFollower"> followers: {{this.followerCount}} </button>
    
            <button class="btn" @click="getFollowing"> following: {{this.followingCount}} </button>
            <button class="btn" @click="getBans"> bans: {{this.bannedCount}} </button>

         
   
            <button class="btn"> photos: {{this.photoCount}} </button>
    </div>
    </div>

    <div id="followerModal" class="modal">


<div class="modal-content">
<span @click="closeModal" class="close">&times;</span>
<div v-for="follow in this.followers" >
<p> {{ follow }}  </p>


</div>
</div>
</div>

<div id="followingModal" class="modal">


<div class="modal-content">
<span @click="closeModal" class="close">&times;</span>
<div v-for="follow in this.following" >
<p> {{ follow }}  </p>


</div>
</div>
</div>

<div id="banModal" class="modal">


<div class="modal-content">
<span @click="closeModal" class="close">&times;</span>
<div v-for="ban in this.bans" >
<p> {{ ban }}  </p>


</div>
</div>
</div>



    <div class="gro">
    <input type="text" id="username" v-model="this.newUser" placeholder="Enter new username" ><br>
    <button class="btn btn-outline-dark" id="submit" @click="updateUsername"> submit  </button>
    </div>

    <br>

    <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css">
 
    </head>

    <div class="container">
        <div class="card-columns">
            <div v-for="photo in this.photos" style="width=300px">
                <div class="card">
                    <img class="card-img-top" :src=photo.photo  alt="unavailable" >
                    <hr>
                    <div class="card-body">
                        <button type="button" class="btn btn-danger" style="float: right;" @click="deletePhoto(photo.photoID)">Delete</button>

                    <button class="fa fa-heart-o" v-if="photo.isLiked==true" @click="this.unLikePhoto(photo.photoID)"> {{ photo.like_count }}</button>
                    <button class="fa fa-heart-o"  v-if="photo.isLiked==false" @click="this.LikePhoto(photo.photoID)"> {{ photo.like_count }}</button>
                    <button  @click="popupComment"> comments</button>



                    <div id="commentModal" class="modal">

                    <!-- Modal content -->
                    <div class="modal-content">
                    <span @click="closeModal" class="close">&times;</span>
                    <div v-for="comment in photo.comment_list" :key="photo.photoID">
                    <p> <b>{{ comment.username }} :</b> {{ comment.comment }} </p>
                    

                    </div>
                    </div>
                </div>

                    </div>

                    <br>
                    <br>
                    <p > likes: {{ photo.like_count }}<br>
                    comments: {{ photo.comment_count }}<br>
                     date: {{ photo.date }}</p>
   
                    <div class="gro">
                    <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="input comment">
                        <div class="input-group-append">
                            <button class="btn btn-outline-dark" type="button" @click="uploadComment(photo.photoID, photo.comment)">post</button>
                        </div>
                    </div><br>
                    </div>
                
                </div>

            </div>
        </div>
</template>


<style>

.modal{

    max-height: 100vh;
    background-color: rgb(179, 226, 226);
    position: absolute;

}

.fa-hearto {
  color: red;
  background-color: grey;
  cursor: pointer;
  height: 27px;
}
.fa-heart {
  color: red;
  cursor: pointer;
  background-color: red;
  height: 27px;
}

.commenter{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:flex-start;
 
}
.card{
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    width: 100px;
    height:800px;
    border-color: blue;  
}



.combtn{
    border-color: black;
    margin-right: 200px;
    width: 50px;
    height:30px;
    

}

.container {
  margin-top: -100px;
  margin-right: 50px;
}

.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

</style>
