package reptiletool

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func CallMd5(data []byte) string {

	md := md5.New()
	md.Write(data)
	return hex.EncodeToString(md.Sum(nil))
}

func CallSha1(data []byte)string{

	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

var CityList = [][]string{
	{"广州市", "113.280637", "23.125178"},
	{"深圳市", "114.085947", "22.547"},
	{"佛山市", "113.122717", "23.028762"},
	{"东莞市", "113.746262", "23.046237"},
	{"珠海市", "113.553986", "22.224979"},
	{"中山市", "113.382391", "22.521113"},
	{"清远市", "113.051227", "23.685022"},
	{"肇庆市", "112.472529", "23.051546"},
	{"惠州市", "114.412599", "23.079404"},
	{"潮州市", "116.632301", "23.661701"},
	{"河源市", "114.697802", "23.746266"},
	{"江门市", "113.094942", "22.590431"},
	{"揭阳市", "116.355733", "23.543778"},
	{"茂名市", "110.919229", "21.659751"},
	{"梅州市", "116.117582", "24.299112"},
	{"汕头市", "116.708463", "23.37102"},
	{"汕尾市", "115.364238", "22.774485"},
	{"韶关市", "113.591544", "24.801322"},
	{"阳江市", "111.975107", "21.859222"},
	{"云浮市", "112.044439", "22.929801"},
	{"湛江市", "110.364977", "21.274898"},
	{"上海市", "121.472644", "31.231706"},
	{"北京市", "116.405285", "39.904989"},
	{"天津市", "117.190182", "39.125596"},
	{"重庆市", "106.504962", "29.533155"},
	{"杭州市", "120.153576", "30.287459"},
	{"宁波市", "121.549792", "29.868388"},
	{"温州市", "120.672111", "28.000575"},
	{"绍兴市", "120.582112", "29.997117"},
	{"舟山市", "122.106863", "30.016028"},
	{"台州市", "121.428599", "28.661378"},
	{"嘉兴市", "120.750865", "30.762653"},
	{"湖州市", "120.102398", "30.867198"},
	{"金华市", "119.649506", "29.089524"},
	{"丽水市", "119.921786", "28.451993"},
	{"衢州市", "118.87263", "28.941708"},
	{"苏州市", "120.619585", "31.299379"},
	{"无锡市", "120.301663", "31.574729"},
	{"南京市", "118.767413", "32.041544"},
	{"徐州市", "117.184811", "34.261792"},
	{"盐城市", "120.139998", "33.377631"},
	{"常州市", "119.946973", "31.772752"},
	{"南通市", "120.864608", "32.016212"},
	{"连云港市", "119.178821", "34.600018"},
	{"扬州市", "119.421003", "32.393159"},
	{"淮安市", "119.021265", "33.597506"},
	{"宿迁市", "118.275162", "33.963008"},
	{"泰州市", "119.915176", "32.484882"},
	{"镇江市", "119.452753", "32.204402"},
	{"厦门市", "118.11022", "24.490474"},
	{"福州市", "119.306239", "26.075302"},
	{"泉州市", "118.589421", "24.908853"},
	{"漳州市", "117.661801", "24.510897"},
	{"莆田市", "119.007558", "25.431011"},
	{"龙岩市", "117.02978", "25.091603"},
	{"南平市", "118.178459", "26.635627"},
	{"宁德市", "119.527082", "26.65924"},
	{"三明市", "117.635001", "26.265444"},
	{"武汉市", "114.298572", "30.584355"},
	{"宜昌市", "111.290843", "30.702636"},
	{"荆州市", "112.23813", "30.326857"},
	{"咸宁市", "114.328963", "29.832798"},
	{"荆门市", "112.204251", "31.03542"},
	{"襄阳市", "112.144146", "32.042426"},
	{"鄂州市", "114.890593", "30.396536"},
	{"恩施土家族苗族自治州", "109.48699", "30.283114"},
	{"黄冈市", "114.879365", "30.447711"},
	{"黄石市", "115.077048", "30.220074"},
	{"潜江市", "112.896866", "30.421215"},
	{"十堰市", "110.787916", "32.646907"},
	{"随州市", "113.37377", "31.717497"},
	{"天门市", "113.165862", "30.653061"},
	{"仙桃市", "113.453974", "30.364953"},
	{"孝感市", "113.926655", "30.926423"},
	{"长沙市", "112.982279", "28.19409"},
	{"郴州市", "113.032067", "25.793589"},
	{"株洲市", "113.151737", "27.835806"},
	{"湘潭市", "112.944052", "27.82973"},
	{"衡阳市", "112.607693", "26.900358"},
	{"岳阳市", "113.132855", "29.37029"},
	{"张家界市", "110.479921", "29.127401"},
	{"益阳市", "112.355042", "28.570066"},
	{"邵阳市", "111.46923", "27.237842"},
	{"常德市", "111.691347", "29.040225"},
	{"怀化市", "109.97824", "27.550082"},
	{"永州市", "111.608019", "26.434516"},
	{"娄底市", "112.008497", "27.728136"},
	{"湘西土家族苗族自治州", "109.739735", "28.314296"},
	{"成都市", "104.065735", "30.659462"},
	{"宜宾市", "104.630825", "28.760189"},
	{"绵阳市", "104.741722", "31.46402"},
	{"德阳市", "104.398651", "31.127991"},
	{"广安市", "106.633369", "30.456398"},
	{"达州市", "107.502262", "31.209484"},
	{"南充市", "106.082974", "30.795281"},
	{"乐山市", "103.761263", "29.582024"},
	{"广元市", "105.829757", "32.433668"},
	{"泸州市", "105.443348", "28.889138"},
	{"攀枝花市", "101.716007", "26.580446"},
	{"资阳市", "104.641917", "30.122211"},
	{"内江市", "105.066138", "29.58708"},
	{"雅安市", "103.001033", "29.987722"},
	{"自贡市", "104.773447", "29.352765"},
	{"眉山市", "103.831788", "30.048318"},
	{"巴中市", "106.753669", "31.858809"},
	{"遂宁市", "105.571331", "30.513311"},
	{"阿坝藏族羌族自治州", "102.221374", "31.899792"},
	{"甘孜藏族自治州", "101.963815", "30.050663"},
	{"凉山彝族自治州", "102.258746", "27.886762"},
	{"海口市", "110.33119", "20.031971"},
	{"三亚市", "109.508268", "18.247872"},
	{"青岛市", "120.355173", "36.082982"},
	{"济南市", "117.000923", "36.675807"},
	{"济宁市", "116.587245", "35.415393"},
	{"烟台市", "121.391382", "37.539297"},
	{"威海市", "122.116394", "37.509691"},
	{"德州市", "116.307428", "37.453968"},
	{"滨州市", "118.016974", "37.383542"},
	{"东营市", "118.66471", "37.434564"},
	{"菏泽市", "115.469381", "35.246531"},
	{"聊城市", "115.980367", "36.456013"},
	{"临沂市", "118.326443", "35.065282"},
	{"日照市", "119.461208", "35.428588"},
	{"泰安市", "117.129063", "36.194968"},
	{"潍坊市", "119.107078", "36.70925"},
	{"枣庄市", "117.557964", "34.856424"},
	{"淄博市", "118.047648", "36.814939"},
	{"太原市", "112.549248", "37.857014"},
	{"大同市", "113.295259", "40.09031"},
	{"长治市", "113.113556", "36.191112"},
	{"晋城市", "112.851274", "35.497553"},
	{"晋中市", "112.736465", "37.696495"},
	{"临汾市", "111.517973", "36.08415"},
	{"吕梁市", "111.134335", "37.524366"},
	{"朔州市", "112.433387", "39.331261"},
	{"忻州市", "112.733538", "38.41769"},
	{"阳泉市", "113.583285", "37.861188"},
	{"运城市", "111.003957", "35.022778"},
	{"鞍山市", "122.995632", "41.110626"},
	{"本溪市", "123.770519", "41.297909"},
	{"朝阳市", "120.451176", "41.576758"},
	{"大连市", "121.618622", "38.91459"},
	{"丹东市", "124.383044", "40.124296"},
	{"抚顺市", "123.921109", "41.875956"},
	{"阜新市", "121.648962", "42.011796"},
	{"葫芦岛市", "120.856394", "40.755572"},
	{"锦州市", "121.135742", "41.119269"},
	{"辽阳市", "123.18152", "41.269402"},
	{"盘锦市", "122.06957", "41.124484"},
	{"沈阳市", "123.429096", "41.796767"},
	{"铁岭市", "123.844279", "42.290585"},
	{"营口市", "122.235151", "40.667432"},
	{"大庆市", "125.11272", "46.590734"},
	{"大兴安岭地区", "124.711526", "52.335262"},
	{"哈尔滨市", "126.642464", "45.756967"},
	{"鹤岗市", "130.277487", "47.332085"},
	{"黑河市", "127.499023", "50.249585"},
	{"鸡西市", "130.975966", "45.300046"},
	{"佳木斯市", "130.361634", "46.809606"},
	{"牡丹江市", "129.618602", "44.582962"},
	{"七台河市", "131.015584", "45.771266"},
	{"齐齐哈尔市", "123.95792", "47.342081"},
	{"双鸭山市", "131.157304", "46.643442"},
	{"绥化市", "126.99293", "46.637393"},
	{"伊春市", "128.899396", "47.724775"},
	{"白城市", "122.841114", "45.619026"},
	{"白山市", "126.427839", "41.942505"},
	{"长春市", "125.3245", "43.886841"},
	{"吉林市", "126.55302", "43.843577"},
	{"辽源市", "125.145349", "42.902692"},
	{"四平市", "124.370785", "43.170344"},
	{"松原市", "124.823608", "45.118243"},
	{"通化市", "125.936501", "41.721177"},
	{"延边朝鲜族自治州", "129.513228", "42.904823"},
	{"石家庄市", "114.502461", "38.045474"},
	{"秦皇岛市", "119.586579", "39.942531"},
	{"保定市", "115.482331", "38.867657"},
	{"沧州市", "116.857461", "38.310582"},
	{"承德市", "117.939152", "40.976204"},
	{"邯郸市", "114.490686", "36.612273"},
	{"衡水市", "115.665993", "37.735097"},
	{"廊坊市", "116.704441", "39.523927"},
	{"唐山市", "118.175393", "39.635113"},
	{"邢台市", "114.508851", "37.0682"},
	{"张家口市", "114.884091", "40.811901"},
	{"郑州市", "113.665412", "34.757975"},
	{"洛阳市", "112.434468", "34.663041"},
	{"开封市", "114.341447", "34.797049"},
	{"许昌市", "113.826063", "34.022956"},
	{"驻马店市", "114.024736", "32.980169"},
	{"安阳市", "114.352482", "36.103442"},
	{"鹤壁市", "114.295444", "35.748236"},
	{"焦作市", "113.238266", "35.23904"},
	{"南阳市", "112.540918", "32.999082"},
	{"平顶山市", "113.307718", "33.735241"},
	{"三门峡市", "111.194099", "34.777338"},
	{"商丘市", "115.650497", "34.437054"},
	{"新乡市", "113.883991", "35.302616"},
	{"信阳市", "114.075031", "32.123274"},
	{"周口市", "114.649653", "33.620357"},
	{"漯河市", "114.026405", "33.575855"},
	{"濮阳市", "115.041299", "35.768234"},
	{"西安市", "108.948024", "34.263161"},
	{"汉中市", "107.028621", "33.077668"},
	{"咸阳市", "108.705117", "34.333439"},
	{"安康市", "109.029273", "32.6903"},
	{"宝鸡市", "107.14487", "34.369315"},
	{"铜川市", "108.979608", "34.916582"},
	{"榆林市", "109.741193", "38.290162"},
	{"商洛市", "109.939776", "33.868319"},
	{"渭南市", "109.502882", "34.499381"},
	{"延安市", "109.49081", "36.596537"},
	{"安庆市", "117.043551", "30.50883"},
	{"蚌埠市", "117.363228", "32.939667"},
	{"巢湖市", "117.874155", "31.600518"},
	{"池州市", "117.489157", "30.656037"},
	{"滁州市", "118.316264", "32.303627"},
	{"阜阳市", "115.819729", "32.896969"},
	{"合肥市", "117.283042", "31.86119"},
	{"淮北市", "116.794664", "33.971707"},
	{"淮南市", "117.018329", "32.647574"},
	{"黄山市", "118.317325", "29.709239"},
	{"六安市", "116.507676", "31.752889"},
	{"马鞍山市", "118.507906", "31.689362"},
	{"宿州市", "116.984084", "33.633891"},
	{"铜陵市", "117.816576", "30.929935"},
	{"芜湖市", "118.376451", "31.326319"},
	{"宣城市", "118.757995", "30.945667"},
	{"亳州市", "115.782939", "33.869338"},
	{"抚州市", "116.358351", "27.98385"},
	{"赣州市", "114.940278", "25.85097"},
	{"高安市", "115.381527", "28.420951"},
	{"吉安市", "114.986373", "27.111699"},
	{"景德镇市", "117.214664", "29.29256"},
	{"九江市", "115.992811", "29.712034"},
	{"南昌市", "115.892151", "28.676493"},
	{"萍乡市", "113.852186", "27.622946"},
	{"上饶市", "117.971185", "28.44442"},
	{"新余市", "114.930835", "27.810834"},
	{"宜春市", "114.391136", "27.8043"},
	{"鹰潭市", "117.033838", "28.238638"},
	{"昆明市", "102.712251", "25.040609"},
	{"丽江市", "100.233026", "26.872108"},
	{"玉溪市", "102.543907", "24.350461"},
	{"临沧市", "100.08697", "23.886567"},
	{"保山市", "99.167133", "25.111802"},
	{"楚雄彝族自治州", "101.546046", "25.041988"},
	{"大理白族自治州", "100.225668", "25.589449"},
	{"德宏傣族景颇族自治州", "98.578363", "24.436694"},
	{"迪庆藏族自治州", "99.706463", "27.826853"},
	{"红河哈尼族彝族自治州", "103.384182", "23.366775"},
	{"怒江傈僳族自治州", "98.854304", "25.850949"},
	{"普洱市", "100.972344", "22.777321"},
	{"曲靖市", "103.797851", "25.501557"},
	{"文山壮族苗族自治州", "104.24401", "23.36951"},
	{"西双版纳傣族自治州", "100.797941", "22.001724"},
	{"昭通市", "103.717216", "27.336999"},
	{"安顺市", "105.932188", "26.245544"},
	{"毕节市", "105.28501", "27.301693"},
	{"贵阳市", "106.713478", "26.578343"},
	{"六盘水市", "104.846743", "26.584643"},
	{"黔东南苗族侗族自治州", "107.977488", "26.583352"},
	{"黔南布依族苗族自治州", "107.517156", "26.258219"},
	{"黔西南布依族苗族自治州", "104.897971", "25.08812"},
	{"铜仁市", "109.191555", "27.718346"},
	{"遵义市", "106.937265", "27.706626"},
	{"白银市", "104.173606", "36.54568"},
	{"定西市", "104.626294", "35.579578"},
	{"甘南藏族自治州", "102.911008", "34.986354"},
	{"嘉峪关市", "98.277304", "39.786529"},
	{"金昌市", "102.187888", "38.514238"},
	{"酒泉市", "98.510795", "39.744023"},
	{"兰州市", "103.823557", "36.058039"},
	{"临夏回族自治州", "103.212006", "35.599446"},
	{"陇南市", "104.929379", "33.388598"},
	{"平凉市", "106.684691", "35.54279"},
	{"庆阳市", "107.638372", "35.734218"},
	{"天水市", "105.724998", "34.578529"},
	{"武威市", "102.634697", "37.929996"},
	{"张掖市", "100.455472", "38.932897"},
	{"果洛藏族自治州", "100.242143", "34.4736"},
	{"海北藏族自治州", "100.901059", "36.959435"},
	{"海东市", "102.10327", "36.502916"},
	{"海门市", "121.176609", "31.893528"},
	{"海南藏族自治州", "100.619542", "36.280353"},
	{"海西蒙古族藏族自治州", "97.370785", "37.374663"},
	{"黄南藏族自治州", "102.019988", "35.517744"},
	{"西宁市", "101.778916", "36.623178"},
	{"玉树藏族自治州", "97.008522", "33.004049"},
	{"百色市", "106.616285", "23.897742"},
	{"北海市", "109.119254", "21.473343"},
	{"崇左市", "107.353926", "22.404108"},
	{"防城港市", "108.345478", "21.614631"},
	{"桂林市", "110.299121", "25.274215"},
	{"贵港市", "109.602146", "23.0936"},
	{"河池市", "108.062105", "24.695899"},
	{"贺州市", "111.552056", "24.414141"},
	{"来宾市", "109.229772", "23.733766"},
	{"柳州市", "109.411703", "24.314617"},
	{"南宁市", "108.320004", "22.82402"},
	{"钦州市", "108.624175", "21.967127"},
	{"梧州市", "111.297604", "23.474803"},
	{"玉林市", "110.154393", "22.63136"},
	{"固原市", "106.285241", "36.004561"},
	{"石嘴山市", "106.376173", "39.01333"},
	{"吴忠市", "106.199409", "37.986165"},
	{"银川市", "106.278179", "38.46637"},
	{"中卫市", "105.189568", "37.514951"},
	{"阿克苏地区", "80.265068", "41.170712"},
	{"阿拉尔市", "81.285884", "40.541914"},
	{"阿勒泰地区", "88.13963", "47.848393"},
	{"巴音郭楞蒙古自治州", "86.150969", "41.768552"},
	{"博尔塔拉蒙古自治州", "82.074778", "44.903258"},
	{"昌吉回族自治州", "87.304012", "44.014577"},
	{"哈密市", "93.51316", "42.833248"},
	{"和田地区", "79.92533", "37.110687"},
	{"喀什地区", "75.989138", "39.467664"},
	{"克拉玛依市", "84.873946", "45.595886"},
	{"克孜勒苏柯尔克孜自治州", "76.172825", "39.713431"},
	{"石河子市", "86.041075", "44.305886"},
	{"塔城地区", "82.985732", "46.746301"},
	{"图木舒克市", "79.077978", "39.867316"},
	{"吐鲁番市", "89.184078", "42.947613"},
	{"乌鲁木齐市", "87.617733", "43.792818"},
	{"五家渠市", "87.526884", "44.167401"},
	{"伊犁哈萨克自治州", "81.317946", "43.92186"},
	{"阿拉善盟", "105.706422", "38.844814"},
	{"巴彦淖尔市", "107.416959", "40.757402"},
	{"包头市", "109.840405", "40.658168"},
	{"赤峰市", "118.956806", "42.275317"},
	{"鄂尔多斯市", "109.99029", "39.817179"},
	{"呼和浩特市", "111.670801", "40.818311"},
	{"呼伦贝尔市", "119.758168", "49.215333"},
	{"通辽市", "122.263119", "43.617429"},
	{"乌海市", "106.825563", "39.673734"},
	{"乌兰察布市", "113.114543", "41.034126"},
	{"锡林郭勒盟", "116.090996", "43.944018"},
	{"兴安盟", "122.070317", "46.076268"},
	{"阿里地区", "80.105498", "32.503187"},
	{"昌都市", "97.178452", "31.136875"},
	{"拉萨市", "91.132212", "29.660361"},
	{"林芝市", "94.362348", "29.654693"},
	{"日喀则市", "88.885148", "29.267519"},
	{"山南市", "91.766529", "29.236023"},
	{"香港特别行政区", "114.173355", "22.320048"},
	{"澳门特别行政区", "113.54909", "22.198951"},
}