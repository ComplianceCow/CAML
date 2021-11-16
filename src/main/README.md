

resource pattern:
/api/v1/providers/{providerid}
/api/v1/controls/{controlid}/metrics?active=true
/api/v1/controls/{controlid}/metrics/reports?latest=true
/api/v1/controls/{controlid}/metrics/{metricid}/reports?latest=true

/api/v1/controls/{controlid}/metrics/{metricid}/measures/{measurename}
/api/v1/controls/{controlid}/metrics/{metricid}/measures/reports?latest=true
/api/v1/controls/{controlid}/metrics/{metricid}/measures/{measurename}/reports?latest=true

Parameters for metrics:
 ?active=true

Parameters for reports:
 ?latest=true || ?report_date_from=''&report_date_to=''
 ?provider_name=''
 ?domain=''&org=''&group=''
 

API calls:
1. register provider info
2. set active metric/s
3. list active metrics
4. publish metric report/s
5. fetch metric report/s 
