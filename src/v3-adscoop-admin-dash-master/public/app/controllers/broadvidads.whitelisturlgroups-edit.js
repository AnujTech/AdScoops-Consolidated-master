app.controller('BroadvidAdsWhitelistUrlGroupsEditCtrl', function($scope, $stateParams, $http, $resource) {
  $scope.group = {}

  $scope.uNew = "";
  /*** Anuj **/
  $scope.errorLabelUrl=false;
  $scope.errorLabelName=false;
  $scope.removeUrl = function(idx) {
    var url_to_delete = $scope.group.Urls[idx];

    $scope.group.Urls.splice(idx, 1);
  }

  $scope.addUrl = function() {
    if($scope.uNew=='' || typeof $scope.uNew=='undefined'){
      $scope.errorLabelUrl="Url can not be blanck";
      return false;
    }
    if ($scope.group.Urls === null) {
      $scope.group.Urls = [$scope.uNew];
    } else {
      $scope.group.Urls.push($scope.uNew);
    }
    /*** Anuj **/
    if($scope.group.Urls!=null || typeof $scope.group.Urls!='undefined'){
    $scope.errorLabelUrl=false;
    }
    $scope.uNew = "";
  }

  $scope.saveGroup = function() {

      /*** Anuj
      Url validation
       **/
    var errorCounter=0;
    if($scope.group.Urls===null || typeof $scope.group.Urls=='undefined'){
    $scope.errorLabelUrl="Please input atleast one Url";
    errorCounter++;
    }else if(typeof $scope.group.Urls=='object'){
    if(Object.keys($scope.group.Urls).length==0){
    $scope.errorLabelUrl="Please input atleast one Url";
    errorCounter++;
    }
    }
    if($scope.group.Name=='' || typeof $scope.group.Name=='undefined'){
    $scope.errorLabelName="Please input Name";
    errorCounter++;
    }

    if(errorCounter>0) {
      return false;
    }
    /**** end code ***/


    $http.post("/broadvidads/whitelisturlgroups/save", $scope.group).
    success(function() {
      $scope.errorLabelUrl=false;
      $scope.errorLabelName=false;
    })
  }

  if (typeof $stateParams.id !== 'undefined') {
    $http.get("/broadvidads/whitelisturlgroups/view/" + $stateParams.id).
    success(function(data) {
      $scope.group = data;
    })
  }
})
