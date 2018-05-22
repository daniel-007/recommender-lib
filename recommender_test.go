package recommender

type PostTest struct {
	ID    string `indexable:"id"`
	Title string `indexable:"content"`
	Body  string `indexable:"content"`
}

var posts = []PostTest{
	PostTest{
		// Original post: https://medium.com/@samlansky/the-theory-of-visitors-4c7dd3a1b6d4
		ID:    "ac59ab05-172b-4f09-9b17-1dfe1cb41156",
		Title: "The theory of visitors",
		Body: `I went on a Tinder date with a schoolteacher who read me poetry aloud on his faded leather couch, but he seemed a little too earnest for me, so I went on a Hinge date with a musician who sang songs to me from the piano, but I was afraid to get involved with another creative type, and then I went on a Raya date with a movie producer who took me to a dinner so fancy it felt like a brag, and I imagined what it would be like to be a stepfather to his daughters, flaxen-haired sprites with names like Annabelle or Clarissa, but then I decided I was too young for all that, and I went on a date with a college student I met on Bumble who told me he couldn’t afford to eat out, so we sat on a curb on Sunset eating soft corn tacos from a truck on the corner, and for a moment I felt older than I really was, older than I had ever been before, though in fact I wasn’t even 30 yet. I went on dates with older guys and learned to get their references, the same allusions to movies and television shows released before I was born that seemed to be touchstones for gay men of a certain age — of course I love Beaches! — but I also went on dates with guys my own age or even younger, and I was comfortable with their language, too, Snapchatting selfies from my bed captioned “tired af” dotted with sleepy-eyed emojis. When I went on dates with successful guys, I knew what to say, commiserating over how crowded Soho House had become (it’s overrun!), but later I would complain to friends about their uninterrogated privilege and the high likelihood that they had secret cocaine habits, because rich guys so often do. When I went on dates with guys who were broke, I related to them, too, that needling anxiety of feeling like you never have enough in a city where everyone seems to have so much, but I would rule them out — telling friends that I needed someone more “worldly” and “accomplished,” this politely coded classism that still allowed me to feel good about myself.
		I went on dates in cities all over the country, wherever I was, even if I was only there for a night or two, finding some guy on an app who might keep me company over dinner or drinks. I went on dates when I was happy, and I went on dates when I was sad. I went on dates to feel complete when I felt empty, and when I felt complete on my own I went on dates then, too, because surely, I thought, I should want to be in a partnership composed of two whole people. These motivations were equally powerful — both the allure of fullness when I was starving and the allure of a complement to something that was already just fine on its own, the way a nice wine might pair with a good meal. I went on dates even when I didn’t want to, when I would have preferred to stay home and watch Netflix or go out with my friends, because if I did not go on dates I might never find love, and I knew that love was the highest calling.
		Never mind that I wasn’t even sure whether I was capable of falling in love again, at least with anyone who would love me back. My partner, a man I thought I would marry, left me unexpectedly, and I ended up fleeing New York for Los Angeles, though I still went back to New York sometimes for work and there I went on dates too. I didn’t want to be alone in my new life, free as I was from the shackles of a long-term relationship. I heard people say sometimes that they weren’t dating so they could focus on themselves, and I always found this curious, maybe because my identity as a single man was so slippery. I saw myself with the sharpest focus through the prism of a date, certainly more than I ever did when I was alone, puttering around my apartment or walking to work. On a date, my roughest qualities shifted into relief as I put my best face forward, highlighting only my most appealing features. There, I could be relaxed, dynamic, charismatic, prone to digressions of affectionate self-deprecation that I thought were charming — I hoped they were, at least. Alone, I was anxious. My internal monologue circled a nucleus of self-loathing, recursive and redundant, like a planet orbiting a sun. I never liked myself more than when I was with a guy who liked me.
		But after the date ended, whether with a friendly hug or a lingering kiss or even sex, whatever it was, the moment I was alone again, loneliness would roll in like New England fog. I felt profoundly sad. I had gotten sober years earlier, when I was still a teenager, so I couldn’t medicate the feelings away with wine or pills. Sometimes I would pick up junk food and overeat, looking for satiety that the night had failed to give me, but that bad habit outgrew its usefulness too. Mostly I called old friends to tell them how well things were going or sat alone on my roof, looking out at the city and trying to will the feelings away, swiping endlessly on dating apps and making conversation with strangers.
		One night, not long after I moved to Los Angeles, after having dinner with a handsome but dull young man I’d met on Tinder, I drove up into the hills to my friend Debby’s house. It was the sort of balmy summer night that California did so well, with the smog settling heavy over the Santa Monica Mountains. Once there, I collapsed into a chair on her porch.
		“I’m so sick of going on dates,” I said.
		“Then why do you keep doing it?” Debby asked.
		“Because if I do it enough, then eventually I won’t have to do it anymore,” I said. “But I can’t seem to get anyone to stick around.” I hesitated. “Or maybe I’m the one who never sticks around. Sometimes it’s hard to tell.”
		She leaned in. “Do you believe in the theory of visitors?” She said this conspiratorially, as if she was sharing with me a secret.
		“What’s that?” I asked.`,
	},
	PostTest{
		// Original post: https://medium.com/@francois.chollet/the-impossibility-of-intelligence-explosion-5be4a9eda6ec
		ID:    "dc869e14-0aad-4a3a-b066-6eaf7bd97ec8",
		Title: "The impossibility of intelligence explosion",
		Body: `In 1965, I. J. Good described for the first time the notion of “intelligence explosion”, as it relates to artificial intelligence (AI):
		Let an ultraintelligent machine be defined as a machine that can far surpass all the intellectual activities of any man however clever. Since the design of machines is one of these intellectual activities, an ultraintelligent machine could design even better machines; there would then unquestionably be an “intelligence explosion,” and the intelligence of man would be left far behind. Thus the first ultraintelligent machine is the last invention that man need ever make, provided that the machine is docile enough to tell us how to keep it under control.
		Decades later, the concept of an “intelligence explosion” — leading to the sudden rise of “superintelligence” and the accidental end of the human race — has taken hold in the AI community. Famous business leaders are casting it as a major risk, greater than nuclear war or climate change. Average graduate students in machine learning are endorsing it. In a 2015 email survey targeting AI researchers, 29% of respondents answered that intelligence explosion was “likely” or “highly likely”. A further 21% considered it a serious possibility.
		The basic premise is that, in the near future, a first “seed AI” will be created, with general problem-solving abilities slightly surpassing that of humans. This seed AI would start designing better AIs, initiating a recursive self-improvement loop that would immediately leave human intelligence in the dust, overtaking it by orders of magnitude in a short time. Proponents of this theory also regard intelligence as a kind of superpower, conferring its holders with almost supernatural capabilities to shape their environment — as seen in the science-fiction movie Transcendence (2014), for instance. Superintelligence would thus imply near-omnipotence, and would pose an existential threat to humanity.
		This science-fiction narrative contributes to the dangerously misleading public debate that is ongoing about the risks of AI and the need for AI regulation. In this post, I argue that intelligence explosion is impossible — that the notion of intelligence explosion comes from a profound misunderstanding of both the nature of intelligence and the behavior of recursively self-augmenting systems. I attempt to base my points on concrete observations about intelligent systems and recursive systems.
		The reasoning behind intelligence explosion, like many of the early theories about AI that arose in the 1960s and 1970s, is sophistic: it considers “intelligence” in a completely abstract way, disconnected from its context, and ignores available evidence about both intelligent systems and recursively self-improving systems. It doesn’t have to be that way. We are, after all, on a planet that is literally packed with intelligent systems (including us) and self-improving systems, so we can simply observe them and learn from them to answer the questions at hand, instead of coming up with evidence-free circular reasonings.
		To talk about intelligence and its possible self-improving properties, we should first introduce necessary background and context. What are we talking about when we talk about intelligence? Precisely defining intelligence is in itself a challenge. The intelligence explosion narrative equates intelligence with the general problem-solving ability displayed by individual intelligent agents — by current human brains, or future electronic brains. This is not quite the full picture, so let’s use this definition as a starting point, and expand on it.`,
	},
	PostTest{
		// Original post: https://medium.com/@mokan9997/hidden-in-plain-sight-4761be7b8115
		ID:    "bd1df59b-366a-4afa-90e2-cbf82f2c6bc6",
		Title: "Hidden in Plain Sight",
		Body: `In early 1989, I was working as a business reporter in Washington, D.C., and interviewing a private investigator for a story about his company. At the end of our conversation, he casually mentioned he’d done some research into the assassination of President John F. Kennedy.
		“Oh yeah?” I said, in an ironic, indifferent way. “So who killed him?” He said: “Well, I think this gunsmith in Baltimore, a guy named Howard Donahue, figured it out.” The private eye showed me a magazine article from 1977 and as I read it, my skepticism began to fade. Within a week, I was heading to nearby Towson, Maryland, to meet Donahue in person.
		Donahue’s Dealey Plaza odyssey began 21 years earlier. The World War II veteran was a firearms specialist who’d testified as an expert witness in multiple shooting cases. He was also a well-known marksman. That’s why he’d been recruited to take part in a CBS News reenactment of the shooting in the spring of 1967.
		The network wanted to know if Lee Harvey Oswald’s Italian military-surplus, bolt-action rifle really could have been fired three times with two hits on a moving target in less than six seconds. Donahue proved that it could. Of the 11 shooters participating in the experiment, only Donahue exceeded Oswald’s performance by scoring three hits in 4.8 seconds, well under the 5.6-second maximum.
		The gunsmith’s involvement in the CBS program sparked a stubborn curiosity about President Kennedy’s killing that eventually led him to pursue his own investigation. He zeroed in on the ballistic and forensic evidence and examined witness statements and photographs. He labored in his basement on weekends and evenings and pored over the Warren Commission report and every book he could find. From the start, Donahue was confident his research would support the Warren Commission’s contention that Oswald, acting alone, fired three times: The first shot hit both Kennedy and Texas Governor John Connally, the second missed and the third struck the president in the head.
		But after a decade of work, Donahue reluctantly arrived at a very different conclusion. Based on his assessment of all available evidence, the gunsmith believed the fatal head shot was actually an accident, inadvertently fired by one of the Secret Service agents from the open security car just behind the presidential limousine in the chaotic, final moments of Oswald’s ambush.
		According to Donahue’s analysis, the shooting unfolded like this: Oswald fired his first shot from a sixth-floor window of the Texas School Book Depository soon after the motorcade made the turn onto Elm Street. But his scope was not adjusted properly, according to the Warren Commission, and the bullet missed, hitting the pavement behind and to the right of Kennedy’s limousine. Fragments ricocheted up and struck the inside windshield trim. At least two caught the president in the scalp and caused him to cry out, “My God! I’m hit!”
		Oswald chambered a second round. This time, he skipped the rifle’s offset-mounted scope and instead drew a bead along the iron sights on top of the barrel. He fired again and the bullet ripped into Kennedy’s upper back, exited his neck and pierced Gov. Connally’s right side.
		At this moment, photos from Dealey Plaza show Secret Service Special Agent George W. Hickey Jr. — riding in the left-rear seat of the follow-up car and sitting up high near the trunk deck — already had turned completely around and was looking back toward the book depository. He may well have spotted the barrel of Oswald’s rifle protruding from the sixth-floor window.
		So Hickey reached down and grabbed the Colt AR15 select-fire, semi-automatic rifle from the floor of the car and flipped the safety lever off. He started to stand and turn to acquire Oswald’s position and return fire. But the follow-up car braked suddenly to avoid a collision with the presidential limousine and Hickey lost his balance. His finger slipped off the trigger guard and the weapon discharged. The bullet was flying at 3,300 feet-per-second when it slammed into the back of Kennedy’s head, 21 feet away, and disintegrated.
		Among the many assassination theories that had emerged by 1977 — the year Donahue’s conclusions became public — no one had ever suggested Kennedy’s death may have been due to a friendly fire incident. But before Donahue, no one with any real firearms expertise had independently examined the evidence, and the gunsmith was certain the facts pointed nowhere else. What’s more, he had many friends in law enforcement or with combat experience, and he himself had spent a lifetime around guns. So he understood how easily unintended discharges happened, particularly amid the fear, confusion and adrenaline of a gunfight.
		That the evidence also pointed to a wide-ranging effort by the Secret Service to suppress the truth surprised Donahue not at all. Clearly Oswald had been trying to kill Kennedy. Divulging the actual origin of the head shot would have served no one’s interest. A cover-up, on the other hand, accomplished several objectives and none, it could be argued, were unreasonable or malicious. Hiding the accident safeguarded the nation’s image at the height of the Cold War, preserved the reputation of the Secret Service, protected the legacy of the martyred president and, perhaps most importantly, shielded Agent Hickey and his family from the enduring infamy such a revelation could bring.
		My meeting with Donahue in the winter of 1989 ultimately led me to write a book called Mortal Error, which chronicled the gunsmith’s investigation and findings. A quarter-century has now passed since the book was published. Over that time, not one of Donahue’s central ballistic conclusions has been refuted. On the contrary, a steady accumulation of new information has continued to point to an accidental shot from Hickey. As impossible as it may seem, as difficult as it may be for most to grasp or accept, particularly those with limited knowledge of firearms, the preponderance of evidence continues to converge around this scenario to the exclusion of all others. Whether you choose to recognize it or not, a staggeringly simple answer to one of America’s most tortuous questions has been hidden in plain sight all along.
		Donahue and I became fast friends after that first meeting, and I also got to know his wife, Katie, well. My wife and I spent many weekends at the Donahues’ rowhouse in Towson through the first half of 1989 as Donahue and I began work on the book. I respected Howard for his encyclopedic knowledge of firearms, his marksmanship and his remarkable investigative prowess. But what really held me in awe was his service during the Second World War.
		In 1944–45, Donahue survived 35 combat missions as a B-17 pilot, flying through swarms of killing flak and watching planes disintegrate around him. One incident in particular illustrated the kind of person he was. On a September 1944 mission to bomb a chemical plant in Ludwigshafen, Germany, Donahue was in the co-pilot’s seat when the aircraft commander took a shell fragment in the skull and went berserk. The bomber drifted perilously close to the nearest aircraft in formation as Donahue wrestled with the wounded man. At the last minute, Donahue was able to kick the right rudder to avoid a collision and the top-turret gunner dropped down to help subdue the captain. Two of the plane’s four engines were shot out, but Donahue still managed to limp the aircraft back to England.
		His actions that day saved 18 men, the nine in his crew plus nine in the bomber they nearly hit, and resulted in Donahue being awarded the Distinguished Flying Cross, one of the highest commendations presented by the Army Air Corps. By the time he returned to Baltimore, First Lieutenant Donahue had received numerous other medals and was one of Maryland’s most decorated fliers of the war. He was 22 years old.
		When I asked him what it was like to fly into combat, he said he was scared to death all the time. He said you just had to keep focused and functioning no matter what. And once on the ground, you tried to stay drunk. All of them did, he said. It was the only way to erase, at least momentarily, the skull-crawling anxiety of facing death, day after day after day.
		Donahue was 66 when we met and if the war had ripped a piece out of him, it didn’t show. He was an amicable fellow, gentlemanly in a down-to-earth, unstuffy way, always quick with a friendly comment or joke and entirely without guile or cynicism. He still liked to drink but wasn’t a drunk. He had the coldest beer in town; kept a refrigerator in his basement set at 34 degrees. When we’d go out to dinner, Howard would order his martini with the olives on the side.
		“Don’t want to displace any alcohol,” he’d say with wink.
		I was 31 and had recently moved to the D.C., area. I’d grown up in Topeka, Kansas, spent the last two years of high school in northern New Hampshire after my parents divorced, then worked a string of blue-collar jobs over the next four years as I bounced around the country. In the spring of 1984, I earned a journalism degree from the University of Kansas and went to work for the Kansas City Business Journal. The paper’s parent company acquired the Washington D.C. Business Journal three years later and I was recruited as part of the turn-around team.
		Working with Donahue, I became familiar with the range of evidence he’d developed in support of his friendly fire scenario. But three key ballistic facts stood out:
		First, the fatal bullet’s trajectory was not compatible with a shot from Oswald. The entrance wound in the back of Kennedy’s skull was just to the right of the the crown, or hair whorl, and the exit point was centered in the upper-right portion of the skull. This indicated a bullet path of left-to-right and down at a relatively shallow angle, not right-to-left and sharply down, as would have been the case had Oswald fired the shot. A bullet from the sixth-floor window of the book depository should have exited the center or left side of the president’s face, not the upper right, frontal portion of his skull, based on the position of Kennedy’s head at the moment of impact. The actual trajectory led directly back to Hickey’s position in the left-rear seat of the follow-up car.`,
	},
}

/*func TestGetBinaryRepresentation(t *testing.T) {
	recommender := New(WordBoundary, true, []int{1})

	vocabulary, err := recommender.Vocabulary(posts)
	assert.NoError(t, err)

	_, err = recommender.getBinaryRepresentation(posts[0], vocabulary)
	assert.NoError(t, err)

	//fmt.Println(binaryRepresentation)
}

func TestBinaryRepresentation(t *testing.T) {
	recommender := New(WordBoundary, true, []int{1})

	vocabulary, err := recommender.Vocabulary(posts)
	assert.NoError(t, err)

	_, err = recommender.BinaryRepresentation(posts, vocabulary)
	assert.NoError(t, err)

	//fmt.Println(binaryRepresentation)
}

func TestFrequency(t *testing.T) {
	recommender := New(WordBoundary, true, []int{1})

	words, err := recommender.getWords(posts)
	assert.NoError(t, err)

	freqs := recommender.Frequency(words)

	recommender.CosineSimilarity(freqs, freqs)
}

func BenchmarkNgrams(b *testing.B) {
	recommender := New(WordBoundary, true, []int{2, 3, 4, 5, 6, 7, 8, 9})

	//var resultSet []map[string]uint32
	var err error
	for i := 0; i < b.N; i++ {
		_, err = recommender.ngrams(posts)
		if err != nil {
			panic(err)
		}
		//assert.NoError(b., err)
	}

	for _, ngram := range resultSet {
		for key, ngramInner := range ngram {
			fmt.Printf("%s -> %d\n", key, ngramInner)
		}
	}
}*/
