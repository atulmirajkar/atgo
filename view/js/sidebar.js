

var SidebarBehavior = function() 
{

}

SidebarBehavior.prototype = {

    constructor: SidebarBehavior,
    onLoad: function (){
	window.alert("test");
    }
    
};

function initializeSideBar()
{
    var sidebar = new SidebarBehavior();
    //sidebar.onLoad();
}

document.onload = initializeSideBar();
