
var editFormData;
	
function getFormData() {
	return {
			name:document.getElementById("name").value,
			email:document.getElementById("email").value
	}
}

function clearFormData() {
	document.getElementById("name").value = "";
	document.getElementById("email").value = "";
}
	
function setFormData(name,email) {
	document.getElementById("name").value = name;
		document.getElementById("email").value = email;
}
	
	// set the message for form status
function setSuccessMessage(message) {
	document.getElementById("message").innerHTML = message;
}

function editDataCall(id) {
	// call get user details by id API
	fetch("http://localhost:3000/crud/getUserByID?id="+id,{
		method:"GET"
	}).then((res)=>res.json()).then((response)=>{
		console.log("Edit info",response);
		editFormData =  response[0];
		setFormData(editFormData.name,editFormData.email)
	})
}
	
// callled this function when user click on button
function submitForm() {
	if(!editFormData) addUser(); // if the editFormData is undefined then call addUser()
 	else editData();
}

// add user function 
function addUser() {
	let payload  = getFormData();
	fetch("http://localhost:3000/crud/insertData",{
  	method:"POST",
			headers:{
				"Content-Type":"application/json"
			},
			body:JSON.stringify(payload)
		}).then((res)=>res.json()).then((response)=>{
			setSuccessMessage(response.message)
				// clear input email and name
				clearFormData();
				getData(); // reload table 
		})
}
	
// edit data 
function editData() {
	var formData = getFormData();
	formData['id'] = editFormData._id; // get _id from selected user
	fetch("http://localhost:3000/crud/updateData",{
		method:"POST",
		headers:{
			"Content-Type":"application/json"
		},
		body:JSON.stringify(formData)
	}).then((res)=>res.json()).then((response)=>{
		setSuccessMessage(response.message)
			clearFormData(); // clear the form field
			getData() // reload the table
	})
}
	
// delete data
function deleteData(id) {
	fetch("http://localhost:3000/crud/delete?id="+id).then((res)=>res.json()).then(
		(response)=>{
			setSuccessMessage(response.message);
			getData();
		}
	)
}
	
// get data method
function getData() {
	fetch("http://localhost:3000/api/products").then(
		(res)=>res.json()
			).then((response)=>{
				var tmpData  = "";
				console.log(response);
				response.forEach((products)=>{
					tmpData+="<div class='my-2 col-lg-3 col-md-4 col-sm-6 col-6'>";
  					tmpData+="<div class='d-flex flex-column border border-secondary rounded'>";
							tmpData+="<div style='width: auto;'>";
								tmpData+="<img src='"+products.Image+"' class='card-img-top' alt='"+products.Name+"'>";
							tmpData+="</div>";
							tmpData+="<div class='p-2'>";
									tmpData+="<h5 class='card-title'>"+products.Name+"</h5>";
									tmpData+="<span>"+products.Shop+"</span>";
									tmpData+="<h6 class='card-title'>"+products.Price+"</h6>";
									tmpData+="<a href='/products/detail/"+products.Id+"' class='btn btn-info my-2'>Detail</a>";
									// tmpData+="<td><button class='btn btn-danger' onclick='deleteData(`"+user._id+"`)'>Delete</button></td>";

							tmpData+="</div>";
						tmpData+="</div>";	
					tmpData+="</div>";
							
		})
	  document.getElementById("getData").innerHTML = tmpData;
  })     
}
	
getData();