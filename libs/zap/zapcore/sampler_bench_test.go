package zapcore_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhangdapeng520/zdpgo_log/libs/atomic"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/internal/ztest"
	. "github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

var counterTestCases = [][]string{
	// some stuff I made up
	{
		"foo",
		"bar",
		"baz",
		"alpha",
		"bravo",
		"charlie",
		"delta",
	},

	// shuf -n50 /usr/share/dict/words
	{
		"unbracing",
		"stereotomy",
		"supranervian",
		"moaning",
		"exchangeability",
		"gunyang",
		"sulcation",
		"dariole",
		"archheresy",
		"synchronistically",
		"clips",
		"unsanctioned",
		"Argoan",
		"liparomphalus",
		"layship",
		"Fregatae",
		"microzoology",
		"glaciaria",
		"Frugivora",
		"patterist",
		"Grossulariaceae",
		"lithotint",
		"bargander",
		"opisthographical",
		"cacography",
		"chalkstone",
		"nonsubstantialism",
		"sardonicism",
		"calamiform",
		"lodginghouse",
		"predisposedly",
		"topotypic",
		"broideress",
		"outrange",
		"gingivolabial",
		"monoazo",
		"sparlike",
		"concameration",
		"untoothed",
		"Camorrism",
		"reissuer",
		"soap",
		"palaiotype",
		"countercharm",
		"yellowbird",
		"palterly",
		"writinger",
		"boatfalls",
		"tuglike",
		"underbitten",
	},

	// shuf -n100 /usr/share/dict/words
	{
		"rooty",
		"malcultivation",
		"degrade",
		"pseudoindependent",
		"stillatory",
		"antiseptize",
		"protoamphibian",
		"antiar",
		"Esther",
		"pseudelminth",
		"superfluitance",
		"teallite",
		"disunity",
		"spirignathous",
		"vergency",
		"myliobatid",
		"inosic",
		"overabstemious",
		"patriarchally",
		"foreimagine",
		"coetaneity",
		"hemimellitene",
		"hyperspatial",
		"aulophyte",
		"electropoion",
		"antitrope",
		"Amarantus",
		"smaltine",
		"lighthead",
		"syntonically",
		"incubous",
		"versation",
		"cirsophthalmia",
		"Ulidian",
		"homoeography",
		"Velella",
		"Hecatean",
		"serfage",
		"Spermaphyta",
		"palatoplasty",
		"electroextraction",
		"aconite",
		"avirulence",
		"initiator",
		"besmear",
		"unrecognizably",
		"euphoniousness",
		"balbuties",
		"pascuage",
		"quebracho",
		"Yakala",
		"auriform",
		"sevenbark",
		"superorganism",
		"telesterion",
		"ensand",
		"nagaika",
		"anisuria",
		"etching",
		"soundingly",
		"grumpish",
		"drillmaster",
		"perfumed",
		"dealkylate",
		"anthracitiferous",
		"predefiance",
		"sulphoxylate",
		"freeness",
		"untucking",
		"misworshiper",
		"Nestorianize",
		"nonegoistical",
		"construe",
		"upstroke",
		"teated",
		"nasolachrymal",
		"Mastodontidae",
		"gallows",
		"radioluminescent",
		"uncourtierlike",
		"phasmatrope",
		"Clunisian",
		"drainage",
		"sootless",
		"brachyfacial",
		"antiheroism",
		"irreligionize",
		"ked",
		"unfact",
		"nonprofessed",
		"milady",
		"conjecture",
		"Arctomys",
		"guapilla",
		"Sassenach",
		"emmetrope",
		"rosewort",
		"raphidiferous",
		"pooh",
		"Tyndallize",
	},
}

func BenchmarkSampler_Check(b *testing.B) {
	for _, keys := range counterTestCases {
		b.Run(fmt.Sprintf("%v keys", len(keys)), func(b *testing.B) {
			fac := NewSamplerWithOptions(
				NewCore(
					NewJSONEncoder(testEncoderConfig()),
					&ztest.Discarder{},
					DebugLevel,
				),
				time.Millisecond, 1, 1000)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					ent := Entry{
						Level:   DebugLevel + Level(i%4),
						Message: keys[i],
					}
					_ = fac.Check(ent, nil)
					i++
					if n := len(keys); i >= n {
						i -= n
					}
				}
			})
		})
	}
}

func makeSamplerCountingHook() (func(_ Entry, dec SamplingDecision), *atomic.Int64, *atomic.Int64) {
	droppedCount := new(atomic.Int64)
	sampledCount := new(atomic.Int64)
	h := func(_ Entry, dec SamplingDecision) {
		if dec&LogDropped > 0 {
			droppedCount.Inc()
		} else if dec&LogSampled > 0 {
			sampledCount.Inc()
		}
	}
	return h, droppedCount, sampledCount
}

func BenchmarkSampler_CheckWithHook(b *testing.B) {
	hook, dropped, sampled := makeSamplerCountingHook()
	for _, keys := range counterTestCases {
		b.Run(fmt.Sprintf("%v keys", len(keys)), func(b *testing.B) {
			fac := NewSamplerWithOptions(
				NewCore(
					NewJSONEncoder(testEncoderConfig()),
					&ztest.Discarder{},
					DebugLevel,
				),
				time.Millisecond,
				1,
				1000,
				SamplerHook(hook),
			)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					ent := Entry{
						Level:   DebugLevel + Level(i%4),
						Message: keys[i],
					}
					_ = fac.Check(ent, nil)
					i++
					if n := len(keys); i >= n {
						i -= n
					}
				}
			})
			// We expect to see 1000 dropped messages for every sampled per settings,
			// with a delta due to less 1000 messages getting dropped after initial one
			// is sampled.
			assert.Greater(b, dropped.Load()/1000, sampled.Load()-1000)
			dropped.Store(0)
			sampled.Store(0)
		})
	}
}
