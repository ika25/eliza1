package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"regexp"
	"time"
	"math/rand"
	
	//

)

	
//----------------------------------------------------------------------
type Response struct {
	pattern         *regexp.Regexp
	possibleAnswers []string
}

func substituteWords(ans string, reflectionMap map[string]string) string {
	Words := strings.Split(ans, " ") 
	for index, word := range Words {
		if val, ok := reflectionMap[word]; ok {
			Words[index] = val 
		}
	}
	return strings.Join(Words, " ") 
}

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}


func Eliza(w http.ResponseWriter, r *http.Request){
	t := time.Now()
	

	rand.Seed(time.Now().UnixNano())
	user := r.URL.Query().Get("value")
	
	re := regexp.MustCompile(`(?i)[I|i] [hate|hate](.*)`)

//questions,aswers
	if matched, _:= regexp.MatchString(`(?i).*\bHow|Why\b.*`, user);matched{
		ansq1 := []string{`I like to keep anwers`,`Why would you ask me that?`,`I'm a chatBot`}
		randindex1 := rand.Intn(len(ansq1))
		fmt.Fprintf(w,ansq1[randindex1])
		return
	}
	
		if matched, _:= regexp.MatchString(`(?i).*Hello|Hey|Hi`, user);matched{
						ansq3 := []string{`Hello I'm glad you could drop by today.`,`Hello My name is Eliza!`,`Hi there how are you today?!`}
						randindex3 := rand.Intn(len(ansq3))
						fmt.Fprintf(w,ansq3[randindex3])
							return
					}
					if matched, _:= regexp.MatchString(`(?i).*\bam\b.*`, user);matched{
						ansq2 := []string{`I dont know what that feels like, I am a ChatBotave', always wondered how feelings work.`,`I dont know how that feeling feels`}
						randindex2 := rand.Intn(len(ansq2))
						fmt.Fprintf(w,ansq2[randindex2])
						return
					}

							if matched, _:= regexp.MatchString(`(?i).*hate|dont'|dislike.*`, user);matched{
								ans4 := []string{`Don't be so negative.`,`Try to be happy`,`As a chatBot i try to like everything.`}
								randindex4 := rand.Intn(len(ans4))
								fmt.Fprintf(w,ans4[randindex4])
								return
							}

							if matched, _:= regexp.MatchString(`(?i).*[s|S]orry.*`, user);matched{
											ansq7 := []string{`No need to be sorry , continue`}
											randindex7 := rand.Intn(len(ansq7))
											fmt.Fprintf(w,ansq7[randindex7])
											return
										}

								if matched, _:= regexp.MatchString(`(?i).*[b|B]ecause.*`, user);matched{
												ansq8 := []string{`I'm sorry user sometimes I can't grasp human conversation! Lovely day though?`}
												randindex8 := rand.Intn(len(ansq8))
												fmt.Fprintf(w,ansq8[randindex8])
												return
											}
			

									if matched, _:= regexp.MatchString(`(?i).*day.*`, user);matched{
										ansq5 := []string{`Today is ` + t.Weekday().String()}
										randindex5 := rand.Intn(len(ansq5))
										fmt.Fprintf(w,ansq5[randindex5])
										return
									}
	
									if matched, _:= regexp.MatchString(`(?i).*time.*`, user);matched{
										ansq6 := []string{`The time is ` + t.Format("4:04PM")}
										randindex6 := rand.Intn(len(ansq6))
										fmt.Fprintf(w,ansq6[randindex6])
										return
									}

												if matched, _:= regexp.MatchString(`(?i).*\bfriend\b.*`, user);matched{
													fmt.Fprintf(w,"I am your friend")
													return
												}
	
	
	//change words
	answers := []string{`Why do %s?`, `How long have %s?`,`Why %s?`}
	likeResponse := Response{re, answers} // reads from file.

	//------------------------------------------------------------------------------
	
	reflectionMap := map[string]string{
		"time": t.Format("4:04PM"),
		"I": "you",
		"i": "you",
		"my": "your",
		"your": "my",
		"are": "am",
		"like":"liked",
		"was": "were",
		"hate":"hate",
		
	}

	//------------------------------------------------------------------------------
	if likeResponse.pattern.MatchString(user) {
		match := re.FindStringSubmatch(user)             
		topic := match[0]                                     
		topic = substituteWords(topic, reflectionMap)         
		index := rand.Intn(len(likeResponse.possibleAnswers)) 
		chosenResponse := likeResponse.possibleAnswers[index] 
		finalResponse := fmt.Sprintf(chosenResponse, topic)   
		fmt.Fprintf(w,finalResponse)                            

		// The output will be "How long have you liked me?" or "Why do you like me?"
	} else {
		responses :=[]string{
			"How does that make you feel?",
			"I dont know what you’re trying to say. Could you explain it to me please?",
			"Why do you say that?",
		}
		randindex := rand.Intn(len(responses))
		fmt.Fprintf(w,responses[randindex]) 
	}
	
}

func main() {
	
	
	rand.Seed(time.Now().UnixNano())
	
	//static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/Eliza", Eliza)
	
    log.Println("Enter this in your web browser - Localhost:8088")
    http.ListenAndServe(":8088", nil)	
	
}
