<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>短网址-工具箱-蓦然回首</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!-- Le styles -->
    <link href="/static/css/bootstrap.css" rel="stylesheet">
    <style type="text/css">
      body {
        padding-top: 60px;
        padding-bottom: 40px;
      }
    </style>
    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
  </head>

  <body>
    <div class="container-fluid">
      <div class="row-fluid" style='width:300px;margin:10px auto;'>
        <div class="span3">
		    <fieldset>
				<div class="input-append input-prepend">
					<input type="text" name="url" id='url' value="" />
					<button id='createBtn'  class="btn btn-primary" type='submit'>生成短网址</button>
				</div>
				<div>
				    <br/>
				    <p>
				    短网址： <span id='short_url' style='color:#cc0000;' ></span>
				    </p>
				</div>
		    </fieldset>

        </div><!--/span-->
      </div><!--/row-->

      <hr>

      <footer>
        <p style='text-align:center'>&copy;kisspy 2013</p>
      </footer>

    </div><!--/.fluid-container-->

    <!-- Le javascript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="/static/js/jquery.js"></script>
	<script>
		$('#createBtn').click(function(){
			var url = $.trim($('#url').val());
			if(''==url){
				$('#url').val('').focus();
				alert('请输入网址');
				return false;
			}
			$.post('/',{url:url},function(data){
				if(data.state=='0'){
				    $('#url').focus();
					alert(data.message);
				}
				else{
		            $('#short_url').html(data.short_url);		    
				}
			},'json');
		});
		
	</script>
  </body>
</html>
