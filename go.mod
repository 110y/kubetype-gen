module github.com/110y/kubetype-gen

go 1.16

require (
	cloud.google.com/go/spanner v1.10.0
	cuelang.org/go v0.0.16-0.20200320220106-76252f4b7486
	github.com/client9/gospell v0.0.0-20160306015952-90dfc71015df
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/emicklei/proto v1.6.15
	github.com/getkin/kin-openapi v0.1.1-0.20190507152207-d3180292eead
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/kr/pretty v0.2.0
	github.com/prometheus/alertmanager v0.20.0
	github.com/prometheus/client_golang v1.2.1
	github.com/russross/blackfriday/v2 v2.0.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
	golang.org/x/tools v0.0.0-20210106214847-113979e3529a
	gopkg.in/src-d/go-license-detector.v2 v2.0.0-20180510072912-da552ecf050b
	gopkg.in/yaml.v2 v2.2.8
	istio.io/api v0.0.0-20210107192247-a205c627e4b9
	istio.io/gogo-genproto v0.0.0-20210107204948-697d6f912366
	istio.io/tools v0.0.0-00010101000000-000000000000
	k8s.io/apiextensions-apiserver v0.18.2
	k8s.io/apimachinery v0.20.1
	k8s.io/gengo v0.0.0-20201214224949-b6c5ce23f027
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	sigs.k8s.io/controller-tools v0.4.1
)

replace istio.io/tools => ./
