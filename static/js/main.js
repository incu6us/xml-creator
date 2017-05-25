(function() {
    'use strict';

    angular
        .module('MainApp', [])
        .controller('MainCtrl', MainCtrl);

    MainCtrl.$inject = ['$scope', '$http'];

    function MainCtrl($scope, $http) {
        $scope.createXML = function () {
            var identifier = document.getElementById('identifier').value

            var result = document.getElementById('result')

            var body = {
                "shipment": {
                    "identifier": parseInt(identifier),
                }
            }

            $http({
                method: 'POST',
                url: '/api/create',
                data: body,
            }).then(function ok(response) {
                console.log("ok");
                $scope.result = response.data;
            }, function error(response) {
                alert("backend error");
                $scope.result = response
            });
        }

        $scope.paxCount = 0;
        $scope.addPax = function () {
            $scope.paxCount++;
            var input = document.createElement("input");
            input.type = "text";
            input.id = "pax["+$scope.paxCount+"]";
            document.getElementById('paxes').appendChild(input);
        }
    }
}());