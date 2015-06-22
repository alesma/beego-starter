process.env.BROWSERIFYSHIM_DIAGNOSTICS=1

var $ = require('jquery'), 
  jqVersion = $().jquery; 

$('#jq-version').text(jqVersion);