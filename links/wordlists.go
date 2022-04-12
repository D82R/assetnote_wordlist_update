package links

var automated []string
var technologies []string

func CompileAutomated(date string) []string {
	date = date + ".txt"
	automatedLinks := []string{
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_subdomains_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_directories_1m_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_txt_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_xml_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_cgi_pl_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_js_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_html_htm_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_aspx_asp_cfm_svc_ashx_asmx_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_jsp_jspa_do_action_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_php_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_apiroutes_",
		"https://wordlists-cdn.assetnote.io/data/automated/httparchive_parameters_top_1m_",
	}

	for _, link := range automatedLinks {
		link = link + date
		automated = append(automated, link)
	}
	return automated
}

func CompileTechnologies(date string) []string {
	date = date + ".txt"
	technologyLinks := []string{
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_express_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_django_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_flask_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_cherrypy_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_zend_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_apache_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_nginx_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_tomcat_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_rails_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_laravel_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_yii_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_coldfusion_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_symfony_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_spring_",
		"https://wordlists-cdn.assetnote.io/data/technologies/httparchive_adobe_experience_manager_",
	}

	for _, link := range technologyLinks {
		link = link + date
		technologies = append(technologies, link)
	}
	return technologies
}
