
function createRequest()
{
    var request;
    try{
	request = new XMLHTTPRequest();
    }
    catch(tryMS)
    {
	try{
	    request = new ActiveXObject("Msxml2.XMLHTTP");
	}
	catch(otherMS)
	{
	    try{
		request = new ActiveXObject("Microsoft.XMLHTTP");
	    }
	    catch(failed)
	    {
		request = null;
	    }
	    
	}
    }
    
    return request;
}
