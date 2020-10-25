var updateImage = function () {
$.ajax({
    url: '/ajax',
    type: "post",
    contentType: 'application/json; charset=utf-8',
    data: JSON.stringify({}),
    dataType: 'json',
    success: function(r) {
        $('#tesla').attr("src", "data:image/jpeg;base64," + r.Base64EncodedJPG);
        console.log("success!!");
        console.log(r.Base64EncodedJPG);
    }
});
}
  
updateImage();
