<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			userID:"",

			photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLiked:true,
                comment: "",
				followed_username:"",
                comments:[{
                        commentID:0,
                        commentUser: "",
                        comment:"",

                    }],
                }
            ],
		}
	},
	created(){
		this.userID=localStorage.getItem("userID");

		this.getStream();
		
	},
	methods: {
		async getStream() {

			try {
				console.log("mfs");
				this.userID=localStorage.getItem("userID");


				let response = await this.$axios.get("/user/" + this.userID + "/stream");
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

			} catch (e) {
			}
			this.loading = false;
		},

		async LikePhoto(photoID){

console.log("putting like")

        let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/like");
    console.log(response.data);
    this.getStream();
    
    

},
async unLikePhoto(photoID){

console.log("deleting like")
let response=await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/like");
console.log(response.data);

this.getStream();



},



		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h2 class="h22">My STREAM</h2>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>

	<br>

<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></head>
  
    

<div class="container">
	<div class="card-columns">
		<div v-for="photo in this.photos" style="width=300px">
			<div class="card">
				<RouterLink :to="'user/' + photo.followed_username + '/account'" class="nav-link">
				<img class="card-img-top" :src=photo.photo  alt="unavailable" >
			</RouterLink>
				<hr>
				<div class="card-body">
					<h4>{{ photo.followed_username }}</h4>
				<br>
				<button class="fa fa-heart-o" v-if="photo.isLiked==true" @click="this.unLikePhoto(photo.photoID)" :color="red"> {{ photo.like_count }}</button>
				<button class="fa fa-heart-o"  v-if="photo.isLiked==false" @click="this.LikePhoto(photo.photoID)" :color="green"> {{ photo.like_count }}</button>

				<button  type="button" @click="uploadComment( photo.photoID, photo.comment)">comments</button>

				<br>
				<br>
				<p > likes: {{ photo.like_count }}<br>
				comments: {{ photo.comment_count }}<br>
				 date: {{ photo.date }}</p>

		<div class="input-group mb-3">
				<input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="input comment">
					<div class="input-group-append">
						<button class="btn btn-outline-dark"  type="button" @click="uploadComment(photo.photoID, photo.comment)">post</button>
					</div>
		</div>
				</div>
			
			</div>

		</div>
	</div>
	
	 
 </div>
 

	
</template>

<style>


.row{
    margin-top: 37px;
}
.card{
	padding-block: 10px 200px;
    margin-top: 37px;
}

.fa-heart-o {
  color: rgb(255, 255, 255);
  background-color: rgb(209, 161, 210);
  cursor: pointer;
  height: 27px;
}



.container {
  margin-top: 10px;
  margin-right: 50px;
}












.h22{
  display: flex;
  flex-direction: row;
  align-items:center;
  justify-content: center;

}





</style>
