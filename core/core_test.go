package core

import (
	"github.com/serkas/salias/core/solver"
	"testing"
)

func TestNewClf(t *testing.T) {

	cls := solver.NewNaiveClassifier([]string{"a"})
	if cls.State() != "empty" {
		t.Error("Expected Classifire in `empty` state, got", cls.State())
	}
}

func TestNewClfFromBuilder(t *testing.T) {
	cls := solver.NewNaiveClassifier([]string{"a"})
	if cls.State() != "empty" {
		t.Error("Expected Classifire in `empty` state, got", cls.State())
	}
}

func TestBasicClassifyNaive(t *testing.T) {
	trainTask := []string{"a b", "b c", "a c", "d e", "d d e f"}
	trainSolution := []string{"1", "1", "1", "0", "0"}

	cls := solver.NewNaiveClassifier([]string{"0", "1"})

	_ = cls.TrainOnText(trainTask, trainSolution)

	answer, _ := cls.Solve("a b c")
	if answer != "1" {
		t.Error("Classifier fails on easy task")
	}

	answer2, _ := cls.Solve("d e c")
	if answer2 != "0" {
		t.Error("Classifier fails on easy task")
	}
}

func TestAdvancedClassifyNaive(t *testing.T) {
	trainTask := []string{
		`Experience with JOGL (Java Open GL), Spring, TinkerPop,GUAVA, Trove, JFreeChart, Java Excel API.
		 Experience with social networking algorithms and analysis.
		 Experience with Java Swing development. Experience with the Spring framework.`,

		`Development experience with Core Java.
		Development experience in using selenium framework with Java.
		Please note Development experience is of high priority.
		Exposure to eclipse based development environment
		Excellent Problem-solving and analytical skills
		Creative thinking, willingness and ability to quickly learn new concepts and technologies
		Excellent team player, with strong communication skills, passionate about his/her work, self-motivated.
		Fluency in English`,

		`At least 7 years of related working experience
		Server side application development using Java and J2EE.
		Web GUI development using Java, HTML/AJAX/JavaScript.
		Desktop GUI development using Java, Swing.
		Basic knowledge of Solaris/Linux.`,

		`You have the opportunity to get involved in a hot early stage startup as a contributor and leader
		What You Will Be Doing - Building a strong scalable product across multiple verticals
		Lean software development on a short cycle
		Coding and tackling technical problems
		Coaching and developing team members What You Need for this Position 2 to 3 Years of experience and knowledge of:
		Ruby on Rails JavaScript Lead Engineer Experience AngularJS`,

		`Write quality software in Ruby
		Work with our Product team to design and architect compelling new products/features
		Maintain, improve, and migrate legacy code
		To (quickly) learn anything and everything you can to help the team
		In-depth knowledge of Ruby
		A demonstrated ability to write quality software
		Basic knowledge of complexity analysis and architecture
		An understanding of REST`,

		`INNOVECS is building a team for a new product/platform development and looking for Senior Software Engineer
		(Ruby) to work on the full time basis in our Kiev office.
		Product Description
		The product is a browser add-on that provides user with highly relevant deals and offers while user shops online.
		The widget is displayed on the website you’re currently browsing to provide user with exclusive shopping offers.
		This saves user time & money, helping to get the most out of user online shopping experience.
		Gives user the lowest prices on the best retail sites like Amazon, eBay & more.
		Offers user coupons, deals, rebates, club discounts & everyday bargains.
		Only appears when user need it, staying out of the way when user is not shopping online.
		Is safe & free for all members
		Requirements
		3+ years RoR practical experience
		5+ years overall development experience
		Good architectural skills
		Excellent leading skills
		Good DB knowledge
		Practical experience in 3rd party integration`,
	}
	trainSolution := []string{"java", "java", "java", "ruby", "ruby", "ruby"}

	cls := solver.NewNaiveClassifier([]string{"java", "ruby"})

	_ = cls.TrainOnText(trainTask, trainSolution)

	rubyTask := `3+ years of proven working experience in developing applications with Ruby on Rails
	(Rails 3/4, Ruby 2.0.0 and above);
	Experience with related technologies (JavaScript, PostgreSQL, Redis, Git, AWS, Passenger, Sidekiq, Unicorn);
	CSS, HTML5, SLIM, JQuery UI/UX;
	Git – the ability to merge task into branch’s;
	Knowledge of Linux, MacOS X, or other *nix environments;
	Ability to work in a team;
	Spoken English.`

	javaTask := `BS (MS preferred) in Computer Science or related field
	Strong Java dev skills
	Ability and willingness to quickly understand architecture and code by different teams
	Strong communication and negotiation skills
	Knowledge of any of the following is a strong advantage: Java EE, Spring, scripting languages, databases, Linux, profiling tools
	Experience with customer collaboration and/or working in a large distributed team is also a definite advantage.
	Problem-solving and analytical skills
	Creative thinking, willingness and ability to quickly learn new concepts and technologies
	Excellent team player, passionate about his/her work, self-motivated and driven
	Fluency in English`

	answer, _ := cls.Solve(rubyTask)
	if answer != "ruby" {
		t.Error("Classifier fails to recognize ruby vacancy")
	}

	answer2, _ := cls.Solve(javaTask)
	if answer2 != "java" {
		t.Error("Classifier fails to recognize java vacancy")
	}
}
