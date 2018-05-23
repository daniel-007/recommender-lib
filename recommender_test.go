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
	PostTest{
		// Original post: https://medium.com/@kentkroeger3/how-the-democrats-could-still-blow-the-2018-midterms-aa24dc7f50f9
		ID:    "cef0dae7-6927-4875-977a-39c378394930",
		Title: "How the Democrats could still lose the 2018 midterms",
		Body: `The chances are still good that the Democrats will win control of at least the U.S. House this November.
		The prediction market PredictIt.com gives the Democrats a 69 percent chance of gaining control of the U.S. House. And while the same market gives the Democrats only a 38 percent chance of controlling the U.S. Senate, those are odds the Democrats only a year ago would have made them giddy with optimism.
		My own midterm election models (which you can access here) indicate, under current conditions, the Republicans will lose 37 House seats and 4 Senate seats, putting both chambers under the control of the Democrats.
		The smart money remains solidly in the Democratic Party’s corner for the 2018 midterms.
		Yet, with the 2016 presidential election as a vivid reminder, Democrats know there is no such thing as a ‘sure thing’ in American politics. Despite a year and a half media obsession over the Trump-Russia collusion investigation, on November 7th, the Democrats could still find themselves in the minority in both congressional chambers.
		Here is why that is still a real possibility…
		The Mueller investigation may not end with a definitive conclusion
		For 18 months the Trump-Russia collusion story has dominated the national news media. According to the Tyndall Report, among the Top 20 news stories of 2017, Russia-related stories accounted for 31 percent of the coverage by the three broadcast news networks. The ratio for the cable news networks undoubtedly would be even higher.
		Yet, after all this hostile coverage of the Trump presidency, Democrats need to prepare themselves for this real possibility: The Robert Mueller-led investigation probably will not indict Donald Trump.
		“Mueller will not indict Trump for obstruction of justice or for any other crime. Period. Full stop. End of story. Speculations to the contrary are just fantasy,” wrote Paul Rosenzweig in The Atlantic last January. “Mueller will not indict the president, even for money laundering. The resolution of the current American crisis is going to be political, not criminal. The future lies with Congress and, ultimately, the electorate, not with prosecutors and the courts.”
		Rosenzweig isn’t coming out of nowhere with that conclusion. Twenty years ago he served as a Senior Counsel in the investigation of President Clinton. What is his rationale? “The Department of Justice has a long-standing legal opinion that sitting presidents may not be indicted,” writes Rosenzweig. “First issued in 1973 during the Nixon era, the policy was reaffirmed in 2000, during the Clinton era.”
		And what can we expect from the Mueller probe with respect to President Trump? According to Rosenzweig, he will file a report on his findings with the deputy attorney general, Rod Rosenstein (since the Attorney General Jeff Sessions has recused himself).
		We are going to get a report. And based on what has been revealed so far, the chances of the Mueller probe proving Trump (or anyone in his orbit) colluded with the Russians to impact the 2016 election is less than certain.
		Such an ambiguous outcome might further energize #TheResistance’s resolve to put the Democrats in charge of Congress. But it also has the potential to torpedo the movement’s organizing assumption that Trump is not just a misogynistic buffoon, but a traitor. The reality that Trump is not a Russian tool will be deflating for many of his critics.
		Americans may come to the defense of an embattled president
		Democrats should be concerned about another potential outcome of the Mueller probe. The American public may not just grow weary of the Trump-Russia collusion story, a significant percentage may rally around the president, particularly if the economy remains strong and real progress towards peace is achieved on the Korean peninsula.
		There is an historical precedent for this phenomenon. Kenneth Starr’s probe into an array of controversies surrounding President Bill Clinton, including his relationship with Monica Lewinsky, resulted in the House impeaching the president on Dec. 19, 1998. Interestingly, as evidenced in the graph below, Clinton’s Gallup job approval ratings rose consistently until the House impeachment. And though his ratings did dip during the Senate trial (in which Clinton was not convicted), they never dropped below 50 percent.`,
	},
	PostTest{
		// Original post: https://medium.com/@jimmysong/why-blockchain-is-hard-60416ea4c5c
		ID:    "e81b01f1-10cf-47d5-aa4f-a8d061b1af4a",
		Title: "Why Blockchain is Hard",
		Body: `The hype around blockchain is massive. To hear the blockchain hype train tell it, blockchain will now:
		Solve income inequality
		Make all data secure forever
		Make everything much more efficient and trustless
		Save dying babies
		What the heck is a blockchain, anyway? And can it really do all these things? Can blockchain bring something amazing to industries as diverse as health care, finance, supply chain management and music rights?
		And doesn’t being for Bitcoin mean that you’re pro-blockchain? How can you be for Bitcoin but say anything bad about the technology behind it?
		In this article, I seek to answer a lot of these questions by looking at what a blockchain is and more importantly, what it’s not.
		What is a blockchain?
		To examine some of these claims, we have to define what a blockchain is and herein lies a lot of the confusion. Many companies use the word “blockchain” to mean some sort of magical device by which all their data will never be wrong. Such a device, of course, does not exist, at least when the real world is involved.
		So what is a blockchain? Technically speaking, a blockchain is a linked list of blocks and a block is a group of ordered transactions. If you didn’t understand the last sentence, you can think of a blockchain as a subset of a database, with a few additional properties.
		The main thing distinguishing a blockchain from a normal database is that there are specific rules about how to put data into the database. That is, it cannot conflict with some other data that’s already in the database (consistent), it’s append-only (immutable), and the data itself is locked to an owner (ownable), it’s replicable and available. Finally, everyone agrees on what the state of the things in the database are (canonical) without a central party (decentralized).
		It is this last point that really is the holy grail of blockchain. Decentralization is very attractive because it implies there is no single point of failure. That is, no single authority will be able to take away your asset or change “history” to suit their needs. This immutable audit trail where you don’t have to trust anyone is the benefit that everyone that’s playing with this technology is looking for. This benefit, however, come at a great cost.
		The Cost of Blockchains
		The immutable audit trail uncontrolled by any single party is certainly useful, but there are many costs to create such a system. Let’s examine some of the issues.
		Development is stricter and slower
		Creating a provably consistent system is not an easy task. A small bug could corrupt the entire database or cause some databases to be different than other ones. Of course, a corrupted or split database no longer has any consistency guarantees. Furthermore, all such systems have to be designed from the outset to be consistent. There is no “move fast and break things” in a blockchain. If you break things, you lose consistency and the blockchain becomes corrupted and worthless.
		You may be thinking, why can’t you just fix the database or start over and move on? That would be easy enough to do in a centralized system, but this is very difficult in a decentralized one. You need consensus, or the agreement of all players in the system, in order to change the database. The blockchain has to be a public resource that’s not under the control of a single entity (decentralized, remember?), or the entire effort is a very expensive way to create a slow, centralized database.
		Incentive structures are difficult to design
		Adding the right incentive structures and making sure that all actors in the system cannot abuse or corrupt the database is likewise a large consideration. A blockchain may be consistent, but that’s not very useful if it’s got a lot of frivolous, useless data in it because the costs of putting data into it are very low. Neither is a consistent blockchain useful if it has almost no data because the costs of putting data into it are very high.
		What gives the data finality? How can you ensure that the rewards are aligned with the network goals? Why do nodes keep or update the data and what makes them choose one piece of data over another when they are in conflict? These are all incentive questions that need good answers and they need to be aligned not just at the beginning but at all points in the future as technology and companies change, otherwise the blockchain is not useful.
		Again, you may be wondering why you can’t “fix” some broken incentive. Once again, this is easy in a centralized system, but in a decentralized one, you simply cannot change anything without consensus. There’s no “fixing” anything unless there’s agreement from everyone.
		Maintenance is very costly
		A traditional centralized database only needs to be written to once. A blockchain needs to be written to thousands of times. A traditional centralized database needs to only checks the data once. A blockchain needs to check the data thousands of times. A traditional centralized database needs to transmit the data for storage only once. A blockchain needs to transmit the data thousands of times.
		The costs of maintaining a blockchain are orders of magnitude higher and the cost needs to be justified by utility. Most applications looking for some of the properties stated earlier like consistency and reliability can get such things for a whole lot cheaper utilizing integrity checks, receipts and backups.
		Users are sovereign
		This can be really good as companies don’t like the liability of having user data in the first place. This can be bad, however, if the user is “misbehaving”. There’s no way to kick out the user that’s spamming your blockchain with frivolous data or has figured out a way to profit in some fashion that causes other users lots of inconvenience. This is related to the above observation that incentive structures have to be designed really, really well in that a user that figures out an exploit is not likely to give that up, especially if there’s profit for the user.
		You may be thinking that you can simply refuse service to malicious users, which would be very easy to do in a centralized service. However, unlike a centralized service, refusing service is difficult because no single entity has the authority to kick anyone out. The blockchain has to be impartial and enforce the rules defined by the software. If the rules are insufficient to deter bad behavior, you’re out of luck. There is no “spirit” of the law here. You simply have to deal with malicious or misbehaving actors, possibly for a very long time.
		All upgrades are voluntary
		A forced upgrade is not an option. The other players on the network have no obligation to change to your software. If they did, such a system would be much easier, faster and cheaper to build as a centralized system. The point of a blockchain is that it’s not under the control of a single entity and this is violated with a forced upgrade.`,
	},
	PostTest{
		// Original post: https://medium.com/@cscalfani/goodbye-object-oriented-programming-a59cda4c0e53
		ID:    "7db32b21-4c15-430d-bef9-c139b35d83c0",
		Title: "Goodbye, Object Oriented Programming",
		Body: `I’ve been programming in Object Oriented languages for decades. The first OO language I used was C++ and then Smalltalk and finally .NET and Java.
		I was gung-ho to leverage the benefits of Inheritance, Encapsulation, and Polymorphism. The Three Pillars of the Paradigm.
		I was eager to gain the promise of Reuse and leverage the wisdom gained by those who came before me in this new and exciting landscape.
		I couldn’t contain my excitement at the thought of mapping my real-world objects into their Classes and expected the whole world to fall neatly into place.
		I couldn’t have been more wrong.
		Inheritance, the First Pillar to Fall
		At first glance, Inheritance appears to be the biggest benefit of the Object Oriented Paradigm. All the simplistic examples of shape hierarchies that are paraded out as examples to the newly indoctrinated seem to make logical sense.
		And Reuse is the word of the day. No… make that the year and perhaps evermore.
		I swallowed this whole and rushed out into the world with my newfound insight.
		Banana Monkey Jungle Problem
		With religion in my heart and problems to solve, I started building Class Hierarchies and writing code. And all was right with the world.
		I’ll never forget that day when I was ready to cash in on the promise of Reuse by inheriting from an existing class. This was the moment I had been waiting for.
		A new project came along and I thought back to that Class that I was so fond of in my last project.
		No problem. Reuse to the rescue. All I gotta do is simply grab that Class from the other project and use it.
		Well… actually… not just that Class. We’re gonna need the parent Class. But… But that’s it.
		Ugh… Wait… Looks like we gonna also need the parent’s parent too... And then… We’re going to need ALL of the parents. Okay… Okay… I handle this. No problem.
		And great. Now it won’t compile. Why?? Oh, I see… This object contains this other object. So I’m gonna need that too. No problem.
		Wait… I don’t just need that object. I need the object’s parent and its parent’s parent and so on and so on with every contained object and ALL the parents of what those contain along with their parent’s, parent’s, parent’s…
		Ugh.
		There’s a great quote by Joe Armstrong, the creator of Erlang:
		The problem with object-oriented languages is they’ve got all this implicit environment that they carry around with them. You wanted a banana but what you got was a gorilla holding the banana and the entire jungle.
		Banana Monkey Jungle Solution
		I can tame this problem by not creating hierarchies that are too deep. But if Inheritance is the key to Reuse, then any limits I place on that mechanism will surely limit the benefits of Reuse. Right?
		Right.
		So what’s a poor Object Oriented Programmer, who’s had a healthy helping of the Kool-aid, to do?
		Contain and Delegate. More on this later.
		The Diamond Problem
		Sooner or later, the following problem will rear its ugly and, depending on the language, unsolvable head.
		Most OO languages do not support this, even though this seems to make logical sense. What’s so difficult about supporting this in OO languages?`,
	},
	PostTest{
		// Original post: https://medium.freecodecamp.org/an-introduction-to-nginx-for-developers-62179b6a458f
		ID:    "281df9f7-43fc-4cca-bc54-da21dfa4d86b",
		Title: "An Introduction to NGINX for Developers",
		Body: `Picture this - you’ve created a web application and are now searching for the right web server to host it from.
		Your application might consist of multiple static files — HTML, CSS, and JavaScript, a backend API service or even multiple webservices. Using Nginx might be what you are looking for, and there are couple of reasons for that.
		NGINX is a powerful web server and uses a non-threaded, event-driven architecture that enables it to outperform Apache if configured correctly. It can also do other important things, such as load balancing, HTTP caching, or be used as a reverse proxy.
		NGINX Architecture
		In this article, I’ll cover a few basic steps about how to install and configure the most common parts of NGINX.
		Basic Installation — Architecture
		There are two ways to install NGINX, either using a pre-built binary or building it up from source.
		The first method is much easiest and faster, but building it up from source provides the ability to include various third-party modules that make NGINX even more powerful. It allows us to customize it to fit the needs of the application.
		To install a prebuilt Debian package, the only thing you have to do is:
		sudo apt-get update
		sudo apt-get install nginx
		After the installation process has finished, you can verify everything is OK by running the command below, which should print the latest version of NGINX.
		Your new webserver will be installed at the location /etc/nginx/. If you go inside this folder, you will see several files and folders. The most important ones that will require our attention later are the file nginx.conf and the folder sites-available.
		Configuration Settings
		The core settings of NGINX are in the nginx.conf file, which by default looks like this.
		The file is structured into Contexts. The first one is the events Context, and the second one is the http Context. This structure enables some advanced layering of your configuration as each context can have other nested contexts that inherit everything from their parent but can also override a setting as needed.
		Various things in this file can be tweaked based on your needs, but NGINX is so simple to use that you can go along even with the default settings. Some of the most important pieces of the NGINX config file are:
		worker_processes: This setting defines the number of worker processes that NGINX will use. Because NGINX is single threaded, this number should usually be equal to the number of CPU cores.
		worker_connections: This is the maximum number of simultaneous connections for each worker process and tells our worker processes how many people can simultaneously be served by NGINX. The bigger it is, the more simultaneous users the NGINX will be able to serve.
		access_log & error_log: These are the files that NGINX will use to log any erros and access attempts. These logs are generally reviewed for debugging and troubleshooting.
		gzip: These are the settings for GZIP compression of NGINX responses. Enabling this one along with the various sub-settings that by default are commented out will result in a quite a big performance upgrade. From the sub-settings of GZIP, care should be taken for the gzip_comp_level, which is the level of compression and ranges from 1 to 10. Generally, this value should not be above 6 — the gain in terms of size reduction is insignificant, as it needs a lot more CPU usage. gzip_types is a list of response types that compression will be applied on.
		Your NGINX install can support far more than a single website, and the files that define your server’s sites live in the /etc/nginx/sites-available directory.
		However, the files in this directory aren’t “live” — you can have as many site definition files in here as you want, but NGINX won’t actually do anything with them unless they’re symlinked into the /etc/nginx/sites-enabled directory (you could also copy them there, but symlinking ensures there’s only one copy of each file to keep track of).
		This gives you a method to quickly put websites online and take them offline without having to actually delete any files — when you’re ready for a site to go online, symlink it into sites-enabled and restart NGINX.
		The sites-available directory includes configurations for virtual hosts. This allows the web server to be configured for multiple sites that have separate configurations. The sites within this directory are not live and are only enabled if we create a symbolic link into the sites-enabled folder.
		Either create a new file for you application or edit the default one. A typical configuration looks like the below one.`,
	},
	PostTest{
		// Original post: https://codeburst.io/top-10-javascript-errors-from-1000-projects-and-how-to-avoid-them-2956ce008437
		ID:    "f38ff4e4-185a-42b3-835c-28c430d09b91",
		Title: "Top 10 JavaScript errors from 1000+ projects (and how to avoid them)",
		Body: `To give back to our community of developers, we looked at our database of thousands of projects and found the top 10 errors in JavaScript. We’re going to show you what causes them and how to prevent them from happening. If you avoid these “gotchas,” it’ll make you a better developer.
		Because data is king, we collected, analyzed, and ranked the top 10 JavaScript errors. Rollbar collects all the errors for each project and summarizes how many times each one occurred. We do this by grouping errors according to their fingerprints. Basically, we group two errors if the second one is just a repeat of the first. This gives users a nice overview instead of an overwhelming big dump like you’d see in a log file.
		We focused on the errors most likely to affect you and your users. To do this, we ranked errors by the number of projects experiencing them across different companies. If we looked only at the total number of times each error occurred, then high-volume customers could overwhelm the data set with errors that are not relevant to most readers.
		Here are the top 10 JavaScript errors:
		Each error has been shortened for easier readability. Let’s dive deeper into each one to determine what can cause it and how you can avoid creating it.
		1. Uncaught TypeError: Cannot read property
		If you’re a JavaScript developer, you’ve probably seen this error more than you care to admit. This one occurs in Chrome when you read a property or call a method on an undefined object. You can test this very easily in the Chrome Developer Console.
		This can occur for many reasons, but a common one is improper initialization of state while rendering the UI components. Let’s look at an example of how this can occur in a real-world app. We’ll pick React, but the same principles of improper initialization also apply to Angular, Vue or any other framework.
		There are two important things realize here:
		A component’s state (e.g. this.state) begins life as undefined.
		When you fetch data asynchronously, the component will render at least once before the data is loaded — regardless of whether it’s fetched in the constructor, componentWillMount or componentDidMount. When Quiz first renders, this.state.items is undefined. This, in turn, means ItemList gets items as undefined, and you get an error – "Uncaught TypeError: Cannot read property ‘map’ of undefined" in the console.
		This is easy to fix. The simplest way: Initialize state with reasonable default values in the constructor.
		The exact code in your app might be different, but we hope we’ve given you enough of a clue to either fix or avoid this problem in your app. If not, keep reading because we’ll cover more examples for related errors below.
		2. TypeError: ‘undefined’ is not an object (evaluating
		This is an error that occurs in Safari when you read a property or call a method on an undefined object. You can test this very easily in the Safari Developer Console. This is essentially the same as the above error for Chrome, but Safari uses a different error message.
		3. TypeError: null is not an object (evaluating
		This is an error that occurs in Safari when you read a property or call a method on a null object. You can test this very easily in the Safari Developer Console.
		Interestingly, in JavaScript, null and undefined are not the same, which is why we see two different error messages. Undefined is usually a variable that has not been assigned, while null means the value is blank. To verify they are not equal, try using the strict equality operator:
		One way this error might occur in a real world example is if you try using a DOM element in your JavaScript before the element is loaded. That’s because the DOM API returns null for object references that are blank.
		Any JS code that executes and deals with DOM elements should execute after the DOM elements have been created. JS code is interpreted from top to down as laid out in the HTML. So, if there is a tag before the DOM elements, the JS code within script tag will execute as the browser parses the HTML page. You will get this error if the DOM elements have not been created before loading the script.
		In this example, we can resolve the issue by adding an event listener that will notify us when the page is ready. Once the addEventListener is fired, the init()method can make use of the DOM elements.`,
	},
}

/*func TestFrequency(t *testing.T) {
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
