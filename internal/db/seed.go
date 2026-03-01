package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/hershdoshi55/social/internal/store"
)

var usernames = []string{
	"alice", "bob", "dave", "eve", "frank",
	"grace", "henry", "iris", "jack", "karen",
	"liam", "mia", "noah", "olivia", "peter",
	"quinn", "rose", "sam", "tina", "uma",
	"victor", "wendy", "xander", "yasmine", "zack",
	"amber", "blake", "chloe", "dylan", "ellie",
	"finn", "gwen", "hank", "isla", "jake",
	"kylie", "leo", "maya", "nate", "opal",
	"paige", "ryder", "stella", "troy", "ursula",
	"vince", "willa", "xavier", "yara", "zoe",
}

var titles = []string{
	"Pectus Excavatum and Exercise: What You Need to Know",
	"Top 5 Chest Exercises Safe for Pectus Excavatum Patients",
	"How I Built Strength with Pectus Carinatum: My Fitness Journey",
	"Breathing Techniques to Improve Lung Capacity with Pectus Excavatum",
	"Can Weightlifting Worsen Pectus Excavatum? The Truth",
	"Swimming as the Ideal Sport for Pectus Deformity Correction",
	"Yoga Poses That Help Pectus Excavatum Posture and Flexibility",
	"The Role of Core Strength in Managing Pectus Excavatum",
	"Before and After: Exercise Transformation with Pectus Carinatum",
	"Cardio Training Tips for Athletes Living with Pectus Excavatum",
	"Bracing vs Surgery: How Exercise Fits Into Each Treatment Path",
	"How to Build a Workout Routine Around Pectus Excavatum",
	"Understanding Exercise-Induced Shortness of Breath in Pectus Patients",
	"Resistance Band Workouts Designed for Pectus Excavatum",
	"The Pectus Excavatum Athlete: Competing at a High Level",
	"How Posture Correction Exercises Complement Pectus Treatment",
	"Running with Pectus Excavatum: Tips to Improve Endurance",
	"Shoulder and Back Workouts to Counteract Pectus Posture",
	"Physical Therapy Exercises Recommended After Nuss Procedure Recovery",
	"Does Exercise Alone Fix Pectus Excavatum? What the Research Says",
}

var contents = []string{
	// 0: Pectus Excavatum and Exercise: What You Need to Know
	`Pectus excavatum is the most common chest wall deformity, affecting roughly 1 in 300 people. The sternum caves inward, creating a sunken appearance that can range from mild to severe. For many people the condition is purely cosmetic, but for others it compresses the heart and lungs enough to cause real symptoms during physical activity — shortness of breath, rapid heartbeat, chest tightness, and fatigue that hits earlier than it should.

	If you have pectus excavatum and want to stay active, the good news is that exercise is not only safe for the vast majority of patients, it is actively encouraged. The key is understanding your own severity, knowing which activities tend to aggravate symptoms, and building a routine that works with your anatomy rather than against it.

	Mild cases typically have no functional limitation. You can train however you like. Moderate to severe cases may notice symptoms under high cardiovascular load — sprinting, HIIT, competitive sports with sustained effort. In these situations the heart simply has less room to fill, so cardiac output drops before it should.

	Before starting any new exercise program, get a baseline assessment from a cardiologist or pulmonologist familiar with pectus. A CT scan gives you the Haller Index, the standard measure of severity. An echocardiogram can reveal whether the heart is being compressed. Knowing your numbers takes the guesswork out of training decisions.

	Once cleared, focus on exercises that open the chest rather than compress it further. Pulling movements — rows, pull-ups, face pulls — counteract the forward shoulder posture that pectus often causes. Swimming is widely regarded as the gold standard because the horizontal position reduces cardiac compression and the full-body demand builds both strength and lung capacity. Yoga and pilates address the flexibility and postural components.

	Avoid heavy bench press loaded through a shortened range if it causes pain or pressure. Many pectus patients find that dumbbell pressing with a full stretch is more comfortable than barbell work. Listen to your body and modify rather than push through discomfort.

	Exercise will not structurally fix pectus excavatum. Only surgery or, in adolescents, bracing can do that. But a well-designed fitness routine will strengthen the muscles around the deformity, improve posture, expand functional lung capacity, and make daily life feel significantly better.`,

	// 1: Top 5 Chest Exercises Safe for Pectus Excavatum Patients
	`Choosing the right chest exercises when you have pectus excavatum is less about avoiding the chest entirely and more about being smart with load, range of motion, and positioning. Here are five moves that tend to work well.

	1. Dumbbell Chest Flyes (Low to Moderate Weight)
	Flyes stretch the pectorals through a wide arc without the heavy compressive load of a barbell press. Start light, keep a soft bend in the elbows, and lower to where you feel a comfortable stretch — not where you feel pressure behind the sternum. This movement also encourages external rotation of the shoulders, which helps counter the rounded posture pectus often creates.

	2. Cable Crossovers
	The cable machine keeps constant tension on the chest throughout the movement and lets you adjust the angle easily. High-to-low cables mimic the lower chest fibers; low-to-high target the upper chest. Because you are not pinned under a fixed bar, you can shift your body to find the most comfortable position.

	3. Push-Up Variations
	Classic push-ups are generally well tolerated. Wide-grip push-ups emphasize the chest; incline push-ups reduce the load while still building the pushing pattern. If standard push-ups cause sternum discomfort, elevate your hands on a bench or countertop until you build enough strength and confidence to progress.

	4. Seated Chest Press Machine
	Machine presses offer a fixed path of motion, which some pectus patients find easier to control than free weights. The seat back provides support and the handles can usually be adjusted to a comfortable grip width.

	5. Resistance Band Presses and Flyes
	Bands are the most pectus-friendly option of all. The resistance increases as you push or pull away from the anchor, meaning the hardest part of the movement is at end-range where your anatomy is least compressed. Great for warming up, deload weeks, or anyone just returning to training.

	What to avoid or modify: Heavy barbell bench press loaded to failure, exercises that cause sharp sternum pain, and anything that leaves you noticeably more breathless than the effort warrants. When in doubt, reduce the weight and check in with a sports medicine doctor.`,

	// 2: How I Built Strength with Pectus Carinatum: My Fitness Journey
	`I was diagnosed with pectus carinatum at fourteen. My sternum pushed outward, giving my chest a pronounced keel-shaped appearance. Kids noticed. I noticed. I spent most of high school hiding under baggy shirts and avoiding anything that required taking mine off.

	When I started at the gym at seventeen I had no idea what I was doing. I just knew I wanted to feel less self-conscious. My first few months were a disaster of ego-lifting and poor programming, but something unexpected happened: the more I trained, the less I thought about my chest. The muscle I put on around the deformity made it far less noticeable, and more importantly, I stopped caring as much.

	Here is what actually worked for me over five years of consistent training.

	Building the back was the single most impactful thing I did for how my chest looked and how I stood. Rows, pull-downs, face pulls, rear delt flyes — I did them all, every session. A thick, wide back creates the illusion of a more proportional chest and naturally pulls your shoulders back, which changes how the carinatum presents.

	For the chest itself I leaned into higher-rep, moderate-weight training. Three sets of fifteen to twenty on cable flyes and dumbbell presses gave me good muscle development without the joint stress of maxing out on bench. I never felt like the carinatum limited my pressing strength — it was more of a cosmetic concern than a functional one in my case.

	Swimming twice a week was a game-changer for my conditioning. The resistance of the water, the breathing patterns, the full-body engagement — after six months my endurance in every other sport improved dramatically.

	Five years on, I am not embarrassed by my chest anymore. The pectus is still there. It always will be. But so are lats I am proud of, and a back that makes my posture something I think about with confidence rather than shame. If you are at the start of this journey, keep going. It gets better.`,

	// 3: Breathing Techniques to Improve Lung Capacity with Pectus Excavatum
	`One of the most overlooked aspects of living with pectus excavatum is breathing mechanics. The inward depression of the sternum physically limits how far the chest wall can expand, which means many pectus patients unconsciously become shallow breathers over time. This creates a cascade of problems: reduced oxygen delivery, early fatigue during exercise, and a tendency toward anxiety driven partly by chronically under-oxygenated breathing patterns.

	The good news is that breathing is a trainable skill. You cannot change bone with breath work, but you can dramatically improve how efficiently you use the lung capacity you have.

	Diaphragmatic Breathing
	The diaphragm is the primary muscle of respiration and most people, pectus or not, never use it fully. Lie on your back, place one hand on your chest and one on your belly. Inhale slowly through the nose and direct the breath into your belly — the lower hand should rise while the upper hand stays relatively still. Exhale fully through pursed lips. Practice five to ten minutes daily. Over weeks this retrains your default breathing pattern.

	Box Breathing
	Inhale for four counts, hold for four, exhale for four, hold for four. This technique is used by military and athletes alike to manage stress responses and improve breath control. For pectus patients it is useful because it forces a full, deliberate expansion of the chest rather than the shallow sips many of us default to.

	Pursed Lip Breathing
	Exhale slowly through pursed lips as if blowing out a candle gently. This creates slight back pressure that keeps airways open longer, similar to how a CPAP works. Useful during exercise when you feel breathless — slow down, purse the lips, and let the exhale fully empty the lungs before the next inhale.

	Inspiratory Muscle Training
	Devices like the Threshold IMT or PowerBreathe provide resistance on the inhale, strengthening the inspiratory muscles over weeks of training. Multiple studies show improvements in respiratory muscle strength and exercise capacity in patients with various restrictive conditions. Worth discussing with a respiratory therapist.

	None of these techniques replaces surgical correction in severe cases, but for mild to moderate pectus they can make a meaningful difference in how you feel during workouts and everyday life.`,

	// 4: Can Weightlifting Worsen Pectus Excavatum? The Truth
	`This is one of the most common questions in pectus communities and the answer is nuanced: for the vast majority of patients, weightlifting will not worsen pectus excavatum structurally. The deformity is determined by abnormal cartilage growth, not by muscle imbalance or exercise habits.

	That said, certain patterns of training can make the cosmetic appearance and the postural consequences worse over time, even if the underlying bone is unchanged.

	The chest-dominant, back-neglected physique is a real phenomenon in gyms everywhere. When you bench press far more than you row, the pectorals become tight and the upper back stays weak. This pulls the shoulders forward and rounds the thoracic spine — a posture that amplifies the sunken appearance of pectus excavatum significantly. From the outside it can look like the condition has gotten worse when really it is posture that has deteriorated.

	Heavy, constant barbell bench pressing in a mechanically disadvantaged position can also cause anterior shoulder and sternum discomfort in some pectus patients, particularly those with moderate to severe grades. This is not damage to the deformity itself, but it is a quality-of-life issue worth taking seriously.

	The prescription is straightforward: balance your training. For every pressing exercise you do, match it with at least one pulling exercise. Prioritize posture-corrective work — face pulls, band pull-aparts, thoracic extensions, wall slides. Keep your pressing volume moderate and use a range of motion that feels comfortable, not one that causes sternum pressure.

	Weightlifting done well is one of the best things you can do for pectus. Strong muscles around the deformity improve appearance, reduce fatigue, and support better posture. The goal is balanced, intelligent training — not avoidance.`,

	// 5: Swimming as the Ideal Sport for Pectus Deformity Correction
	`Ask almost any specialist what sport they recommend for patients with pectus excavatum or carinatum and the answer is nearly always the same: swimming. There are good reasons for this consensus, and whether you are dealing with a functional limitation or purely cosmetic concerns, the pool is worth your time.

	The horizontal position changes everything. When you swim, gravity pulls differently on the chest wall. The heart is no longer fighting the depression to fill properly. Many pectus patients who feel winded quickly on land report significantly better endurance in the water, which makes swimming uniquely accessible for those whose symptoms limit other activities.

	The breathing demands of swimming are extraordinary. Every stroke requires a controlled exhale underwater followed by a quick, full inhale at the surface. Over months and years of practice, this trains the respiratory muscles and the lungs to operate more efficiently within whatever space the chest wall allows. Swimmers routinely develop the largest lung capacities of any sport, and while pectus limits the ceiling, the gains are real and meaningful.

	Freestyle and backstroke are generally the most comfortable for pectus patients. Backstroke in particular opens the chest with every pull cycle, counteracting the forward collapse posture. Butterfly is technically demanding and should wait until you have built a solid foundation, but it is a spectacular developer of the entire upper body.

	Beyond the physiological benefits, swimming is a full-body resistance workout. Water provides about 800 times the resistance of air. Your shoulders, lats, core, and legs all work hard in every session. The result over time is a physique that is wider in the back and shoulders, which visually de-emphasizes the chest deformity considerably.

	If you cannot swim yet, lessons are worth every penny. Even basic lap swimming three times a week for six months will produce noticeable changes in how you breathe, how you look, and how you feel.`,

	// 6: Yoga Poses That Help Pectus Excavatum Posture and Flexibility
	`Yoga is not a cure for pectus excavatum, but it is one of the most effective tools available for addressing the postural and flexibility consequences of the condition. The forward-collapsed posture, the tight chest, the weak upper back — yoga targets all of it.

	Camel Pose (Ustrasana)
	Kneeling with your hands on your heels and your chest lifted toward the ceiling, camel pose creates a deep opening across the front of the chest and shoulders. It directly counteracts the forward curve that pectus reinforces. Hold for five breaths and come out slowly.

	Cobra Pose (Bhujangasana)
	A gentler backbend accessible to all levels. Lying face down, press your hands into the mat and lift your chest. Keep your elbows slightly bent and your shoulders away from your ears. This strengthens the spinal extensors while opening the anterior chest.

	Fish Pose (Matsyasana)
	Lying on your back with your chest lifted and the crown of your head on the mat, fish pose creates a sustained chest opener that many pectus patients find immediately relieving. Support your lower back with your hands if needed and breathe deeply into the front of the chest.

	Thread the Needle
	From a tabletop position, slide one arm under your body and rest your shoulder and ear on the mat. This thoracic rotation stretch releases the upper back tension that builds when the chest is perpetually tight. Alternate sides and hold each for thirty seconds.

	Downward Facing Dog
	A whole-body pose that lengthens the spine, opens the shoulders, and builds the posterior shoulder strength that pectus patients need. Walk your feet in and out to find your ideal position and breathe into the back of the ribcage.

	A daily fifteen-minute yoga routine combining these poses will, over time, visibly improve your posture and create more ease in your chest. Pair it with strength training and cardio for a comprehensive approach.`,

	// 7: The Role of Core Strength in Managing Pectus Excavatum
	`Core strength and pectus excavatum might seem like an unlikely pairing, but the relationship between a strong trunk and a better quality of life with pectus is well established among sports medicine practitioners.

	The core is not just the abs. It is the entire cylinder of musculature that stabilizes the spine and pelvis — the rectus abdominis, obliques, transverse abdominis, multifidus, pelvic floor, and diaphragm. In pectus patients, the diaphragm is often mechanically compromised by the sternal depression, which means it cannot descend as freely as it should during inhalation. A strong surrounding core compensates partially for this limitation.

	Posture is the other critical link. Weak core musculature allows the thoracic spine to round forward, which exaggerates the sunken appearance of pectus excavatum and further restricts the already-limited chest expansion. Building core endurance — not just strength — keeps the spine upright throughout the day, which is when posture actually matters.

	The best core exercises for pectus patients are those that emphasize anti-extension and anti-rotation stability rather than spinal flexion. Dead bugs, pallof presses, bird dogs, and plank variations build the deep stabilizing muscles without reinforcing the forward-flexed posture that crunches and sit-ups tend to create.

	Breathing should be integrated into every core exercise. During a dead bug, exhale fully as you lower the opposite arm and leg, engaging the deep core. During a plank, breathe in a slow, controlled rhythm rather than holding your breath. This trains the core and the respiratory system simultaneously, addressing both of the key functional concerns in pectus.

	Consistency over months rather than intensity in any single session is what produces results. Fifteen minutes of targeted core work four days a week will do more for a pectus patient than occasional high-intensity sessions.`,

	// 8: Before and After: Exercise Transformation with Pectus Carinatum
	`I want to share something honest: exercise did not fix my pectus carinatum. My sternum still protrudes. I still notice it when I look in the mirror without a shirt. But the before and after of five years of dedicated training is still one of the most dramatic changes I have ever experienced, and it happened in ways I did not expect.

	Before: I avoided the beach, public pools, and any situation that might require removing my shirt. My posture was terrible — shoulders hunched forward, head jutting out, chest caved in to minimize the protrusion. I was thin and self-conscious and convinced the condition defined me.

	The first year of training was mostly just showing up. I did not know what I was doing but I went to the gym four times a week and tried things. By month six I had put on about eight pounds of muscle. The carinatum was not gone but suddenly there were muscles around it. My back had some width. My shoulders had some roundness. The protrusion still existed but it was no longer the first thing the eye went to.

	Year two I got intentional. I researched the best exercises for building the aesthetic traits that de-emphasize the deformity — wide lats, thick traps, full rear deltoids. I swam twice a week and added a yoga session. My posture, which had always been terrible, started to genuinely improve because I had the back strength to actually hold it.

	By year five I was competing in a local open-water swim event, something seventeen-year-old me would have laughed at as impossible. I still have carinatum. I am no longer defined by it.

	The transformation was not structural. It was muscular, postural, and psychological. That turned out to be enough.`,

	// 9: Cardio Training Tips for Athletes Living with Pectus Excavatum
	`Cardiovascular training with pectus excavatum requires some adjustments to get the most out of your sessions while staying within your symptomatic limits. These tips are aimed at people who want to push their fitness, not just maintain it.

	Know your limit before you hit it. Pectus excavatum can compress the heart enough to reduce stroke volume, meaning your heart pumps less blood per beat than it should at high intensities. This is why many pectus patients hit an aerobic ceiling earlier than their fitness level would otherwise predict. Training at moderate intensities — 65 to 80 percent of max heart rate — is often more productive than pushing into the zone where symptoms emerge.

	Zone 2 training is your best friend. Long, steady-state efforts at a conversational pace build aerobic base without taxing the cardiovascular system to its limit. Think forty-five minutes on the bike, a long swim, or a slow jog. Over months, zone 2 work raises your aerobic capacity at every intensity level.

	Interval training works, but shorten the hard intervals. Instead of four-minute VO2max efforts, try thirty-second to one-minute hard intervals with full recovery. This lets you accumulate quality cardiovascular stimulus without prolonged time at the intensity where symptoms appear.

	Swimming and cycling are generally better tolerated than running because the former reduces gravitational load on the chest and the latter eliminates the impact. If running is your sport, build mileage slowly and pay attention to the intensities where you feel chest tightness versus where you feel fine.

	Nasal breathing during lower-intensity cardio sessions is worth experimenting with. It forces a slower, more controlled breathing rhythm and improves carbon dioxide tolerance over time, which can reduce the sensation of breathlessness.

	Track your sessions with a heart rate monitor. Data over time will show you your true aerobic ceiling and whether your training is actually moving it upward.`,

	// 10: Bracing vs Surgery: How Exercise Fits Into Each Treatment Path
	`For patients and families navigating the decision between bracing and surgery for pectus deformity, exercise is an important but often underdiscussed component of both treatment paths.

	Bracing, primarily used for pectus carinatum in adolescents whose chest walls are still malleable, applies constant external pressure to gradually remodel the cartilage. The brace is typically worn twenty-three hours a day and removed only for bathing and exercise. That exercise hour is not optional — it is essential. The muscles around the chest need to stay strong and active during a period when they are being compressed all day. Recommended activities during bracing include swimming, cycling, and resistance training with an emphasis on the back and core. Contact sports and anything with a high fall risk are usually restricted to protect the brace and the patient.

	Surgery — the Nuss procedure for excavatum, open repair for carinatum — has a different exercise timeline. The immediate post-operative period, typically the first four to six weeks, is near-complete rest. The sternum is healing and the bar (in Nuss) is settling into position. Gentle walking begins early. After six to eight weeks, light resistance training of the lower body is usually cleared. Upper body work returns gradually over months, with full clearance often not coming until six months post-op.

	What stays constant across both paths is the value of being fit going into treatment. Patients who have strong cardiovascular fitness and good muscular development tend to recover faster from surgery and tolerate bracing better. Think of pre-treatment exercise as prehabilitation.

	Post-treatment exercise is where the real work begins. Once the chest wall has been corrected or is being corrected, building the musculature to support and complement the new shape becomes the priority.`,

	// 11: How to Build a Workout Routine Around Pectus Excavatum
	`Building a sustainable workout routine when you have pectus excavatum is about working within your anatomy intelligently rather than fighting it. Here is a practical framework.

	Assess first. Before writing a single program, understand your severity and symptoms. Can you run a mile without chest tightness? Do you get winded climbing stairs? The answers shape how aggressive your cardio programming should be. If you have not had a medical evaluation, get one. A Haller Index above 3.25 is considered severe and warrants a conversation with a cardiologist before high-intensity training.

	Structure your week around three priorities: cardiovascular health, posterior chain and back development, and posture correction. Everything else is secondary.

	A sample four-day split:
	Monday — Push and posture: incline dumbbell press, cable flyes, face pulls, band pull-aparts, wall slides.
	Wednesday — Cardio: forty-five minutes of swimming or Zone 2 cycling.
	Friday — Pull and core: pull-ups or lat pulldown, seated rows, single-arm dumbbell rows, dead bugs, bird dogs.
	Saturday — Full body and flexibility: goblet squats, Romanian deadlifts, yoga or mobility work.

	Pressing movements should never exceed pulling movements in volume. A 1:1.5 ratio of push to pull is a good rule of thumb. This directly counteracts the posture problems pectus creates.

	Progression should be slow and deliberate. Add weight only when the current load feels truly easy and your form is solid. Chasing numbers on a barbell is less useful than consistently showing up for years.

	Rest and recovery matter more than most people acknowledge. Sleep, nutrition, and stress management are training variables. Shortchange them and your results stall regardless of how well you program.

	Review and adjust every eight to twelve weeks. Your routine should evolve as your fitness, symptoms, and goals change.`,

	// 12: Understanding Exercise-Induced Shortness of Breath in Pectus Patients
	`Shortness of breath during exercise is one of the most common and frustrating symptoms reported by people with moderate to severe pectus excavatum. Understanding why it happens is the first step toward managing it effectively.

	The mechanism is primarily cardiac. When the sternum is depressed, it can push directly against the right ventricle and, in some cases, the right atrium. This compression reduces the chamber's ability to fill with blood during diastole, which limits stroke volume. Because stroke volume is limited, the heart must beat faster to maintain cardiac output at any given exercise intensity. The system reaches its ceiling earlier than it should, producing the sensation of breathlessness at workloads that should feel manageable.

	There is also a respiratory component. The depressed sternum restricts the forward excursion of the chest wall during inhalation, reducing tidal volume. The body compensates by breathing faster, which is less efficient and more tiring, contributing to the perception of breathlessness.

	Not all breathlessness in pectus patients is caused by the deformity. Deconditioning, anxiety, exercise-induced asthma, and poor breathing technique can all produce similar symptoms. A cardiopulmonary exercise test (CPET) is the definitive way to determine whether your breathlessness is cardiac, pulmonary, or something else entirely.

	Management strategies include training primarily in Zone 2 where symptoms are absent or mild, using nasal breathing to slow the breathing rate, incorporating breath training devices to strengthen respiratory muscles, and working with a cardiologist or pulmonologist to rule out other causes.

	For severe cases where breathlessness is significantly limiting quality of life and confirmed to be cardiac in origin, surgical correction of the deformity is often the only thing that produces meaningful improvement. Exercise management helps, but it has limits.`,

	// 13: Resistance Band Workouts Designed for Pectus Excavatum
	`Resistance bands deserve more credit than they get. For pectus excavatum patients, they offer a unique combination of benefits that make them arguably the safest and most versatile training tool available.

	The resistance profile of bands — lighter at the start of a movement, heavier at the end — means the most challenging point is where your chest is most open and least compressed. Compare this to a barbell bench press where the heaviest load is right at the point of maximum sternal compression. Bands simply fit the anatomy better.

	Here is a complete resistance band routine for pectus patients:

	Warm-Up (5 minutes)
	Band pull-aparts: 3 sets of 20. Hold the band at shoulder width and pull it apart until your arms are fully extended. This activates the rear deltoids and rhomboids immediately.

	Main Circuit
	Banded chest press (anchored at mid-chest height): 3 sets of 15. Step forward for more resistance, back for less.
	Banded row (anchored at mid-chest height): 3 sets of 15. Pull the band to your lower ribs and squeeze the shoulder blades together.
	Banded face pull (anchored at face height): 3 sets of 20. External rotation emphasis — essential for posture.
	Banded pull-apart overhead: 2 sets of 15. Hold the band overhead and pull apart, opening the chest and thoracic spine.
	Banded deadbug: 2 sets of 10 per side. Loop the band around your feet and press your lower back flat while alternating arm and leg extensions.

	Cool Down
	Doorway stretch: anchor the band at shoulder height, grab it, and rotate away to get a chest stretch. Hold 30 seconds each side.

	This full routine takes under thirty minutes and hits every postural priority for pectus patients. Do it three times a week and you will notice a difference in posture within a month.`,

	// 14: The Pectus Excavatum Athlete: Competing at a High Level
	`The idea that pectus excavatum prevents high-level athletic performance is contradicted by the evidence — both anecdotal and medical. Numerous elite athletes have competed at national and international levels with the condition. What separates them is not the absence of pectus symptoms but how they have adapted their training and mindset around it.

	The functional limitation of pectus excavatum is real in moderate to severe cases. The cardiac output ceiling is lower. The respiratory reserve is smaller. These are measurable facts on a CPET. But athletes are extraordinary adapters. The body responds to training stimulus regardless of structural limitations, and the aerobic ceiling, while lower than a structurally normal athlete's, can still be raised substantially through consistent training.

	The athletes who do best are those who find their sport carefully. Events that reward efficiency over pure power output tend to suit pectus athletes well. Long-distance swimming is a common example — the horizontal position, the controlled breathing, the technical nature of the sport all play to pectus strengths. Cycling is another, where cardiac output matters but the position is favorable. Rowing, despite being brutally demanding, is embraced by some pectus athletes because it directly strengthens the pulling muscles that matter most.

	Sports with explosive, short-duration demands are generally better tolerated than sustained high-intensity events. Many pectus athletes thrive in team sports where bursts of effort are followed by recovery.

	The mental component is significant. Every pectus athlete has to work through moments of doubt — breathlessness that others do not seem to feel, fatigue that arrives uninvited. The ones who compete at a high level have learned to distinguish normal athletic discomfort from their symptomatic limit, and they train their minds as deliberately as their bodies.`,

	// 15: How Posture Correction Exercises Complement Pectus Treatment
	`Whether you are living with pectus excavatum conservatively, going through bracing, or recovering from surgery, posture correction exercises are one of the most consistently recommended interventions by physical therapists and orthopedic specialists. Here is why they matter and how to do them.

	Pectus excavatum creates a predictable postural pattern. The shoulders roll forward, the thoracic spine rounds into kyphosis, the head shifts forward of the shoulders, and the chest collapses inward. This posture is partly a physical consequence of the deformity and partly a learned compensation — many patients unconsciously hunch to hide their chest.

	The problem is that this posture reinforces itself. Tight pectorals and anterior shoulder muscles pull the shoulders further forward. Weak rhomboids, lower trapezius, and serratus anterior fail to pull them back. The head-forward position creates neck tension. The rounded thoracic spine reduces rib expansion, compounding whatever respiratory limitation already exists.

	Correcting this pattern requires a two-pronged approach: stretching what is tight and strengthening what is weak.

	For stretching: doorway chest stretches, thoracic extension over a foam roller, and pectoral minor stretches against a wall. Hold each thirty to sixty seconds and do them daily.

	For strengthening: face pulls with external rotation, prone Y-T-W exercises, wall angels, and serratus punches. These exercises feel awkward at first because the muscles they target are often severely underdeveloped. Stick with them.

	The timeline for meaningful postural change is months, not weeks. Tissues adapt slowly. But the cumulative effect of daily posture work over six to twelve months is visible to anyone — better shoulder position, a more upright spine, and a chest that presents more favorably regardless of the underlying deformity.`,

	// 16: Running with Pectus Excavatum: Tips to Improve Endurance
	`Running is the activity pectus excavatum patients most commonly report struggling with. The upright position, the sustained cardiovascular demand, and the rhythmic breathing requirements all interact with the functional limitations of the deformity in challenging ways. But running is absolutely achievable for the majority of pectus patients — it just requires a smarter approach than most training plans suggest.

	Start slower than you think you need to. Most people with pectus who struggle with running are running too fast for their aerobic system to handle. Slow down to a pace where you can hold a conversation. It will feel embarrassingly slow at first. That is fine. This is zone 2 training and it builds the aerobic base that makes faster running possible later.

	Nasal breathing while running is a powerful tool. It forces you to run slow enough to maintain the breath, which keeps you in the aerobic zone, and it improves carbon dioxide tolerance over time — reducing the air hunger sensation that many pectus patients experience. Start with short intervals of nasal breathing during easy runs and extend them over weeks.

	Warm up properly before running. Five minutes of brisk walking followed by dynamic stretching of the hips and lower body primes the cardiovascular system gradually rather than shocking it. Cold-starting into a run is particularly rough for pectus patients.

	Interval running is more effective than steady-state running for building speed without prolonged time at high intensities. Walk for two minutes, jog for one, walk for two — repeat. Over weeks, the jogging intervals get longer and the walking intervals shorter.

	Track your heart rate. If you consistently hit your max heart rate at paces that feel moderate, that is a meaningful signal. Share the data with your doctor. It may indicate that surgical correction would produce functional improvements worth pursuing.`,

	// 17: Shoulder and Back Workouts to Counteract Pectus Posture
	`The most impactful thing most pectus patients can do in the gym is ignore their chest and work their back. This sounds counterintuitive but makes complete sense once you understand the postural mechanics at play.

	Pectus excavatum, and to a lesser extent carinatum, promotes a posture where the shoulders round forward, the scapulae wing away from the spine, and the thoracic spine rounds into kyphosis. This posture is driven by tight, shortened anterior muscles and weak, lengthened posterior muscles. The fix is systematic.

	The exercises that matter most, in order of priority:

	Face Pulls — The single best posture correction exercise. Attach a rope to a cable machine at face height, grip both ends, and pull toward your face while rotating your wrists outward. Trains the rear deltoids, external rotators, and mid-traps simultaneously. Do these every session, 3-4 sets of 15-20 reps.

	Prone Y-T-W Raises — Lie face down on a bench or the floor and lift your arms into Y, T, and W shapes. No weight needed. These isolate the lower trapezius and rear deltoids in ways that cable work cannot fully replicate.

	Barbell or Dumbbell Rows — Build the mid and upper back thickness that creates the wide, strong appearance that visually de-emphasizes the chest deformity. Keep the torso close to parallel and drive the elbows back rather than up.

	Pull-Ups — If you can do them, pull-ups are the king of back exercises. They build the lats, which create the V-taper that makes every chest look better by comparison.

	Wall Angels — Stand against a wall, arms in goalpost position, and slide them overhead while keeping contact with the wall. Deceptively difficult. Corrects scapular movement patterns that years of pectus posture have disrupted.

	Program back work at a 1.5:1 ratio to chest work. Within three months of consistent training you will stand differently.`,

	// 18: Physical Therapy Exercises Recommended After Nuss Procedure Recovery
	`Recovery from the Nuss procedure — the minimally invasive surgery used to correct pectus excavatum — is a gradual, months-long process. Physical therapy plays a critical role in returning patients to full function, and the exercises prescribed are carefully sequenced to match the healing timeline.

	Weeks 1-4: The priority is pain management, basic mobility, and prevention of complications like pneumonia. Patients are encouraged to walk short distances multiple times daily, increasing duration as tolerated. Deep breathing exercises are mandatory — the therapist will teach incentive spirometer use and diaphragmatic breathing to prevent atelectasis. Shoulder gentle range-of-motion movements are performed to prevent stiffening.

	Weeks 4-8: As healing progresses, light lower body exercise is introduced. Stationary cycling with no resistance is generally safe by week six. Upper body movement remains restricted — no lifting above shoulder height, no reaching behind the back, no pushing or pulling heavy objects. Core activation through gentle dead bug progressions begins.

	Months 2-4: Light upper body resistance training returns, beginning with bands and progressing to very light dumbbells. The focus remains on posture correction — face pulls, rows, and scapular stabilization exercises. Swimming is typically cleared around month three with a technique modification: no butterfly, and turns should avoid twisting.

	Months 4-6: Progressive return to normal resistance training. The bar is still in place and the sternum is still consolidating, so contact sports, heavy lifting, and high-impact activities remain restricted. Running is usually cleared by month four.

	After bar removal (typically two to three years post-op): Full activity clearance with no restrictions.

	Compliance with physical therapy timelines matters. Patients who rush back to full training risk complications. Follow your surgical team's protocol precisely.`,

	// 19: Does Exercise Alone Fix Pectus Excavatum? What the Research Says
	`The short answer is no. Exercise alone cannot structurally correct pectus excavatum. The deformity involves abnormal costal cartilage growth and sternal positioning — neither of which responds to muscle strengthening or cardiovascular training. No amount of bench pressing, swimming, or yoga will move bone.

	What the research does show, however, is considerably more nuanced and encouraging.

	Studies on respiratory muscle training in pectus patients demonstrate measurable improvements in inspiratory muscle strength, exercise capacity, and breathlessness scores. The chest wall does not change, but how the patient breathes within it does.

	Posture research consistently shows that targeted strengthening of the posterior chain — the rhomboids, lower trapezius, rear deltoids, and spinal extensors — produces significant changes in thoracic kyphosis and shoulder position over three to six months of training. This does not alter the bony deformity but dramatically changes how it presents cosmetically and can reduce the perception of severity.

	Psychological research on pectus patients is less abundant but points clearly in one direction: patients who engage in regular exercise report significantly better body image, self-esteem, and quality of life than sedentary controls, independent of whether they underwent surgical correction. The mechanism appears to be both direct (improved appearance through muscle development) and indirect (exercise-induced mood and confidence improvements).

	The bottom line for someone deciding between exercise and surgery: exercise is not an alternative to surgery for severe functional cases. But for mild to moderate cases where surgery is not indicated, or as a complement to bracing in adolescents, exercise produces real and meaningful improvements. It just does not do what a Nuss bar or open repair does.

	Pursue both if the situation warrants. They are not mutually exclusive.`,
}

var tags = []string{
	"pectus-excavatum",
	"pectus-carinatum",
	"chest-wall",
	"exercise",
	"fitness",
	"posture",
	"breathing",
	"swimming",
	"recovery",
	"surgery",
	"nuss-procedure",
	"cardio",
	"strength-training",
	"physical-therapy",
	"yoga",
	"core",
	"bracing",
	"lung-capacity",
	"athlete",
	"chest-deformity",
}

var comments = []string{
	"Great post, really helpful!",
	"Thanks for sharing this, learned a lot.",
	"This is exactly what I needed to read today.",
	"Nice post!",
	"Very informative, keep it up!",
	"Bookmarking this for later.",
	"I never knew this, thanks!",
	"This helped me understand so much more.",
	"Really well written, appreciate it.",
	"Thanks, this was very useful!",
	"Awesome content as always.",
	"Shared this with my friend who has the same condition.",
	"Would love to see more posts like this.",
	"Super helpful, thank you!",
	"Finally a post that explains this clearly.",
	"This motivated me to get back to the gym.",
	"Interesting read!",
	"Good stuff, thanks for putting this together.",
	"I relate to this so much.",
	"Following for more content like this!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post:= range posts{
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment:= range comments{
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")

}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.IntN(len(users))]

		idx := i % len(titles)

		numTags := rand.IntN(5) + 1
		postTags := make([]string, numTags)
		for j := 0; j < numTags; j++ {
			postTags[j] = tags[rand.IntN(len(tags))]
		}

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[idx],
			Content: contents[idx],
			Tags:    postTags,
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment{
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID: posts[rand.IntN(len(posts))].ID,
			UserID: users[rand.IntN(len(users))].ID,
			Content: comments[rand.IntN(len(comments))],
		}

	}
	return cms
}