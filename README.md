# testcase

    SRSLY-2:testcase bezigon$ godep save
    SRSLY-2:testcase bezigon$ goapp serve
    INFO     2015-05-27 15:27:24,158 devappserver2.py:745] Skipping SDK update check.
    INFO     2015-05-27 15:27:24,503 api_server.py:190] Starting API server at: http://localhost:62415
    INFO     2015-05-27 15:27:24,526 dispatcher.py:192] Starting module "default" running at: http://localhost:8080
    INFO     2015-05-27 15:27:24,528 admin_server.py:118] Starting admin server at: http://localhost:8000
    ERROR    2015-05-27 15:27:26,172 go_runtime.py:180] Failed to build Go application: (Executed command: /Users/bezigon/google-cloud-sdk/platform/google_appengine/goroot/bin/go-app-builder -app_base /Users/bezigon/go/src/github.com/bezigon/testcase -arch 6 -dynamic -goroot /Users/bezigon/google-cloud-sdk/platform/google_appengine/goroot -nobuild_files ^^$ -unsafe -gopath /Users/bezigon/go -print_extras_hash Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/attribute.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/fragment_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/document_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xpath/util.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/node_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/css/css.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/text.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/document_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xpath/xpath_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xpath/xpath.go Godeps/_workspace/src/github.com/moovweb/gokogiri/mem/mem_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/help/help_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/fragment_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/help/help.go Godeps/_workspace/src/github.com/moovweb/gokogiri/mem/mem.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/document.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/element.go Godeps/_workspace/src/github.com/moovweb/gokogiri/util/util_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/css/css_test.go testcase.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/utils_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/nodeset.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/search_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/gokogiri_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xpath/expression.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xpath/util_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/fragment.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/fragment.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/xpath_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/attribute_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/node_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/utils_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/cdata.go Godeps/_workspace/src/github.com/moovweb/gokogiri/gokogiri.go Godeps/_workspace/src/github.com/moovweb/gokogiri/util/util.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/document.go Godeps/_workspace/src/github.com/moovweb/gokogiri/help/util_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/crash_test.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/comment.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/pi.go Godeps/_workspace/src/github.com/moovweb/gokogiri/xml/node.go Godeps/_workspace/src/github.com/moovweb/gokogiri/html/encoding_test.go)

    2015/05/27 18:27:26 go-app-builder: Failed parsing input: package github.com/moovweb/gokogiri/html required, but all its files were excluded by nobuild_files
