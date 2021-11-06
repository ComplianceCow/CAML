var dir = DriveApp.getFoldersByName("Working Group").next();

function generateControls() {
  var controlsSheet = SpreadsheetApp.getActive().getSheetByName("Controls");
  var controlsData = controlsSheet.getDataRange().getValues();
  var totalJSONOutput = '';
  var controlJSONOutput = '';
  var depControls  = [];

  // generate the header
  totalJSONOutput = totalJSONOutput+'{\n\t"controls" :\n\t[\n'

  // Starting from 1 as 0 is the header in the excel spreadsheet
  // Adjust this value to 0 if there are no headers
  for (var controlLoop = 1; controlLoop < controlsData.length; controlLoop++) {
      depControls = controlsData[controlLoop][4].split(",");
      controlJSONOutput='';
      controlJSONOutput=controlJSONOutput+'{\n';
      controlJSONOutput=controlJSONOutput+'"controlDomain" : '+' '+'"'+controlsData[controlLoop][0]+'"'+',\n';      
      controlJSONOutput=controlJSONOutput+'"controlTitle" : '+' '+'"'+controlsData[controlLoop][1]+'"'+',\n';     
      controlJSONOutput=controlJSONOutput+'"controlID" : '+' '+'"'+controlsData[controlLoop][2]+'"'+',\n';    
      controlJSONOutput=controlJSONOutput+'"controlSpec" : '+' '+'"'+controlsData[controlLoop][3]+'"'+',\n';    
      controlJSONOutput=controlJSONOutput+'"dependentControls" : '+'[\n';      
      for (var depControlLoop = 0; depControlLoop < depControls.length; depControlLoop++ ) {
        controlJSONOutput=controlJSONOutput+'\t'+'"'+depControls[depControlLoop].trim()+'"';
        if (depControlLoop == depControls.length-1) {
          controlJSONOutput=controlJSONOutput+'\n'
        } else {
          controlJSONOutput=controlJSONOutput+',\n'
        }     
      }
      controlJSONOutput=controlJSONOutput+'\t\t]\n';
      controlJSONOutput=controlJSONOutput+'},\n';
      totalJSONOutput=totalJSONOutput+controlJSONOutput;
  }

  // generate the footer
  totalJSONOutput = totalJSONOutput+']\n}\n'

  var file = dir.createFile("ccm_controls.json", totalJSONOutput);
  // Logger.log(totalJSONOutput);
}


function generateMetrics() {
  var metricSheet = SpreadsheetApp.getActive().getSheetByName("Metrics");
  var metricData = metricSheet.getDataRange().getValues();
  var totalJSONOutput = '';
  var controlJSONOutput = '';

  // generate the header
  totalJSONOutput = totalJSONOutput+'{\n\t"metrics" :\n\t[\n'


  for (var metricLoop = 1; metricLoop < metricData.length; metricLoop++) {
    var controlID=metricData[metricLoop][2];
    var metricID=metricData[metricLoop][3];
    var metricDescription=metricData[metricLoop][6].escapeChars();
    var metricExpression=metricData[metricLoop][7].escapeChars();
    var metricRule=metricData[metricLoop][8].escapeChars();
    var metricImplementationGuidelines=metricData[metricLoop][9].escapeChars();
    var metricSLODescription=String(metricData[metricLoop][10]).escapeChars();

    controlJSONOutput='';
    controlJSONOutput=controlJSONOutput+'{\n';
    controlJSONOutput=controlJSONOutput+'"'+"controlID"+'": '+'"'+controlID+'",\n';
    controlJSONOutput=controlJSONOutput+'"'+"metricID"+'": '+'"'+metricID+'",\n';
    controlJSONOutput=controlJSONOutput+'"'+"metricDescription"+'": '+'"'+metricDescription+'",\n';
    controlJSONOutput=controlJSONOutput+'"'+"metricExpression"+'": '+'"'+metricExpression+'",\n';
    controlJSONOutput=controlJSONOutput+'"'+"metricRule"+'": '+'"'+metricRule+'",\n';
    controlJSONOutput=controlJSONOutput+'"'+"metricImplementationGuidelines"+'": '+'"'+metricImplementationGuidelines+'",\n';
    controlJSONOutput=controlJSONOutput+'"'+"metricSLODescription"+'": '+'"'+metricSLODescription+'"\n';
    controlJSONOutput=controlJSONOutput+'},\n';
    totalJSONOutput=totalJSONOutput+controlJSONOutput;
  }

  // generate the footer
  totalJSONOutput = totalJSONOutput+']\n}\n';

  var file = dir.createFile("ccm_metrics.json", totalJSONOutput);
  // Logger.log(totalJSONOutput);

}

String.prototype.escapeChars = function() {
    return this.trim().replace(/\"/g, '\\"')
                      .replace(/\n/g, '\\n')
                      .replace(/\t/g, '\\t');
  };

