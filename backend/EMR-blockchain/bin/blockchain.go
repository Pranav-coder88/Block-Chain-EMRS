package emrBlockChain

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"time"
)

// Block contains data that will be written to the blockchain.
type Block struct {
	Position  int
	Data      Transaction
	Timestamp string
	Hash      string
	PrevHash  string
}

// Transaction contains data for a checked out MedicalRecord
type Transaction struct {
	WalletAddress string `json:"wallet_address"`
	UserID        string `json:"user_id"`
	UserRole      string `json:"user_role"`
	UpdatedKey    string `json:"updated_key"`
	UpdatedValue  string `json:"updated_value"`
	IsGenesis     bool   `json:"is_genesis"`
}

// User contains data about user and role
type User struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

// MedicalRecord is the Wallet, containing data about patient's medical record
//type MedicalRecord struct {
//	WalletAddress string   `json:"wallet_address"`
//	FullName      string   `json:"full_name"`
//	Operations    []string `json:"operations"`
//	Prescriptions []string `json:"prescriptions"`
//	Allergies     []string `json:"allergies"`
//	CreationDate  string   `json:"creation_date"`
//}

type GenderEnum bool

const (
	GENDER_MALE   GenderEnum = iota != 0
	GENDER_FEMALE GenderEnum = iota != 0
)

type EducationEnum uint8

const (
	EDUCATION_HIGHSCHOOL   EducationEnum = iota
	EDUCATION_COLLEGE      EducationEnum = iota
	EDUCATION_COLLEGE_GRAD EducationEnum = iota
	EDUCATION_ADVANCED_DEG EducationEnum = iota
)

type MarriageEnum uint8

const (
	MARRIAGE_NIL       MarriageEnum = iota
	MARRIAGE_MARRIED   MarriageEnum = iota
	MARRIAGE_DIVORCED  MarriageEnum = iota
	MARRIAGE_SEPARATED MarriageEnum = iota
	MARRIAGE_WIDOWED   MarriageEnum = iota
	MARRIAGE_PARTNERED MarriageEnum = iota
)

type WorklessEnum uint8

const (
	WORKLESS_RETIRED  WorklessEnum = iota
	WORKLESS_DISABLED WorklessEnum = iota
	WORKLESS_SICK     WorklessEnum = iota
)

type Date struct {
	Year  int `json:"Year"`
	Month int `json:"Month"`
	Day   int `json:"Day"`
}

type DrugInfo struct {
	Name     string `json:"Name"`
	Dose     string `json:"Dose"`
	Duration string `json:"Duration"`
}

type FamilyMember struct {
	Age                  uint8  `json:"Age"`
	HealthAndPsychiatric string `json:"HealthAndPsychiatric"`
	AgeAtDeath           uint8  `json:"AgeAtDeath"`
	Cause                string `json:"Cause"`
}

type PreviousSymptoms struct {
	// General
	RecentWeightGain string `json:"RecentWeightGain"`
	Fatigue          bool   `json:"Fatigue"`
	Weakness         bool   `json:"Weakness"`
	Fever            bool   `json:"Fever"`
	NightSweats      bool   `json:"NightSweats "`

	// Muscle/Joints/Bones
	Numbness       bool   `json:"Numbness"`
	JointPain      bool   `json:"JointPain"`
	MuscleWeakness bool   `json:"MuscleWeakness"`
	JointSwelling  string `json:"JointSwelling"`

	// Ears
	RingingInTheEars bool `json:"RingingInTheEars"`
	LossOfHearing    bool `json:"LossOfHearing"`

	// Eye
	EyePain       bool `json:"EyePain"`
	EyeRedness    bool `json:"EyeRedness"`
	LossOfVision  bool `json:"LossOfVision"`
	BlurredVision bool `json:"BlurredVision"`
	EyeDryness    bool `json:"EyeDryness"`

	// Throat
	FrequentSoreThroats    bool `json:"FrequentSoreThroats"`
	ThroatHoarseness       bool `json:"ThroatHoarseness"`
	DifficultyInSwallowing bool `json:"DifficultyInSwallowing"`
	PainInJaw              bool `json:"PainInJaw "`

	// Heart and Lungs
	ChestPain         bool `json:"ChestPain"`
	Palpitations      bool `json:"Palpitations"`
	ShortnessOfBreath bool `json:"ShortnessOfBreath"`
	Fainting          bool `json:"Fainting"`
	SwollenLegsOrFeet bool `json:"SwollenLegsOrFeet"`
	Cough             bool `json:"Cough"`

	// Nervous System
	Headaches           bool `json:"Headaches"`
	Dizziness           bool `json:"Dizziness"`
	LossOfConsciousness bool `json:"LossOfConsciousness"`
	Tingling            bool `json:"Tingling"`
	MemoryLoss          bool `json:"MemoryLoss "`

	// Stomach and intestines
	Nausea                 bool `json:"Nausea"`
	Heartburn              bool `json:"Heartburn"`
	StomachPain            bool `json:"StomachPain"`
	Vomiting               bool `json:"Vomiting"`
	YellowJaundice         bool `json:"YellowJaundice"`
	IncreasingConstipation bool `json:"IncreasingConstipation"`
	PersistentDiarrhea     bool `json:"PersistentDiarrhea"`
	BloodInStools          bool `json:"BloodInStools"`
	BlackStools            bool `json:"BlackStools"`

	// Skin
	SkinRedness  bool `json:"SkinRedness"`
	Rash         bool `json:"Rash"`
	Bumps        bool `json:"Bumps"`
	HairLoss     bool `json:"HairLoss"`
	ColorChanges bool `json:"ColorChanges"`

	// Blood
	Anemia bool `json:"Anemia"`
	Clots  bool `json:"Clots"`

	// Kidney/Urine/Bladder
	FrequentUrination bool `json:"FrequentUrination"`
	BloodInUrine      bool `json:"BloodInUrine"`

	// Women Only
	AbnormalPapSmear       bool `json:"AbnormalPapSmear"`
	IrregularPeriods       bool `json:"IrregularPeriods"`
	BleedingBetweenPeriods bool `json:"BleedingBetweenPeriods"`
	Pms                    bool `json:"Pms"`

	// Psychiatric
	Depression                    bool `json:"Depression"`
	ExcessiveWorries              bool `json:"ExcessiveWorries"`
	DifficultyFallingAsleep       bool `json:"DifficultyFallingAsleep"`
	DifficultyStayingAsleep       bool `json:"DifficultyStayingAsleep"`
	DifficultiesWithSexualArousal bool `json:"DifficultiesWithSexualArousal"`
	PoorAppetite                  bool `json:"PoorAppetite"`
	FoodCravings                  bool `json:"FoodCravings"`
	FrequentCrying                bool `json:"FrequentCrying"`
	Sensitivity                   bool `json:"Sensitivity"`
	SuicidalThoughts              bool `json:"SuicidalThoughts"`
	Stress                        bool `json:"Stress"`
	Irritability                  bool `json:"Irritability"`
	PoorConcentration             bool `json:"PoorConcentration"`
	RacingThoughts                bool `json:"RacingThoughts"`
	Hallucinations                bool `json:"Hallucinations"`
	RapidSpeech                   bool `json:"RapidSpeech"`
	GuiltyThoughts                bool `json:"GuiltyThoughts"`
	Paranoia                      bool `json:"Paranoia"`
	MoodSwings                    bool `json:"MoodSwings"`
	Anxiety                       bool `json:"Anxiety"`
	RiskyBehavior                 bool `json:"RiskyBehavior"`

	OtherProblems string `json:"OtherProblems"`
}

type WomensReproductiveHistory struct {
	AgeOfFirstPeriod     uint8 `json:"AgeOfFirstPeriod"`
	NumberOfPregancies   uint8 `json:"NumberOfPregancies"`
	NumberOfMiscarriages uint8 `json:"NumberOfMiscarriages"`
	NumberOfAbortions    uint8 `json:"NumberOfAbortions"`
	MenopauseAge         uint8 `json:"MenopauseAge"`
	RegularPeriods       bool  `json:"RegularPeriods"`
}

type SubstanceInfo struct {
	Category           string `json:"Category"`
	AgeOfFirstUsed     uint8  `json:"AgeOfFirstUsed"`
	AmountAndFrequency string `json:"AmountAndFrequency"`
	Duration           string `json:"Duration"`
	LastUsage          string `json:"LastUsage"`
	CurrentlyUsing     bool   `json:"CurrentlyUsing"`
}

type PersonalData struct {
	CreatedDate           time.Time `json:"CreatedDate"`
	Name                  string    `json:"Name"`
	Birthdate             string    `json:"Birthdate"`
	Age                   uint8     `json:"Age"`
	Gender                string    `json:"Gender"`
	ModeOfReach           string    `json:"ModeOfReach"`
	SymptomsBrief         string    `json:"SymptomsBrief"`
	PrevPractitioners     string    `json:"PrevPractitioners"`
	PsychHospitalizations string    `json:"PsychHospitalizations"`
	StatusECT             string    `json:"StatusECT"`
	StatusPsychotherapy   string    `json:"StatusPsychotherapy"`
}

type CurrentMedications struct {
	DrugAllergies []string   `json:"DrugAllergies"`
	Medications   []DrugInfo `json:"Medications"`
}

type PastMedicalHistory struct {
	Diabetes          bool     `json:"Diabetes"`
	HighBloodPressure bool     `json:"HighBloodPressure"`
	HighCholesterol   bool     `json:"HighCholesterol"`
	Hypothyroidism    bool     `json:"Hypothyroidism"`
	Goiter            bool     `json:"Goiter"`
	CancerType        string   `json:"CancerType"`
	Leukemia          bool     `json:"Leukemia"`
	Psoriasis         bool     `json:"Psoriasis"`
	Angina            bool     `json:"Angina"`
	HeartProblems     bool     `json:"HeartProblems"`
	HeartMurmur       bool     `json:"HeartMurmur"`
	Pneumonia         bool     `json:"Pneumonia"`
	PulmonaryEmbolism bool     `json:"PulmonaryEmbolism"`
	Asthma            bool     `json:"Asthma"`
	Emphysema         bool     `json:"Emphysema"`
	Stroke            bool     `json:"Stroke"`
	Epilepsy          bool     `json:"Epilepsy"`
	Cataracts         bool     `json:"Cataracts"`
	KidneyDisease     bool     `json:"KidneyDisease"`
	KidneyStones      bool     `json:"KidneyStones"`
	CrohnsDisease     bool     `json:"CrohnsDisease"`
	Colitis           bool     `json:"Colitis"`
	Anemia            bool     `json:"Anemia"`
	Jaundice          bool     `json:"Jaundice"`
	Hepatitis         bool     `json:"Hepatitis"`
	PepticUlcer       bool     `json:"PepticUlcer"`
	RheumaticFever    bool     `json:"RheumaticFever"`
	Tuberculosis      bool     `json:"Tuberculosis"`
	Aids              bool     `json:"Aids"`
	Others            []string `json:"Others"`
}

type PersonalHistory struct {
	BirthProblems    string `json:"BirthProblems"`
	PlaceOfBirth     string `json:"PlaceOfBirth"`
	MaritalStatus    string `json:"MaritalStatus"`
	LatestOccupation string `json:"LatestOccupation"`
	StatusWorking    bool   `json:"StatusWorking"`
	HoursPerWeek     string `json:"HoursPerWeek"`
	StatusSSI        bool   `json:"StatusSSI"`
	DescSSI          string `json:"DescSSI"`
	LegalProblems    string `json:"LegalProblems"`
	Religion         string `json:"Religion"`
}

type FamilyHistory struct {
	Father                 FamilyMember   `json:"Father"`
	Mother                 FamilyMember   `json:"Mother"`
	Siblings               []FamilyMember `json:"Siblings"`
	Children               []FamilyMember `json:"Children"`
	MaternalRelativeIssues string         `json:"MaternalRelativeIssues"`
	PaternalRelativeIssues string         `json:"PaternalRelativeIssues"`
}

type SystemsReview struct {
	PreviousSymptoms          PreviousSymptoms          `json:"PreviousSymptoms"`
	WomensReproductiveHistory WomensReproductiveHistory `json:"WomensReproductiveHistory"` // need to implement/update in medicalHistory.tsx
}

type SubstanceUse struct {
	Alcohol          SubstanceInfo   `json:"Alcohol"`
	Cannabis         SubstanceInfo   `json:"Cannabis"`
	StimulantsA      SubstanceInfo   `json:"StimulantsA"`
	StimulantsB      SubstanceInfo   `json:"StimulantsB"`
	Amphetamines     SubstanceInfo   `json:"Amphetamines"`
	Tranquilizers    SubstanceInfo   `json:"Tranquilizers"`
	Sedatives        SubstanceInfo   `json:"Sedatives"`
	Heroin           SubstanceInfo   `json:"Heroin"`
	IllicitMethadone SubstanceInfo   `json:"IllicitMethadone"`
	OtherOpioids     SubstanceInfo   `json:"OtherOpioids"`
	Hallucinogens    SubstanceInfo   `json:"OtherOpioids"`
	Inhalants        SubstanceInfo   `json:"Hallucinogens"`
	Others           []SubstanceInfo `json:"Inhalants"`
}

type MedicalRecord struct {
	WalletAddress      string             `json:"wallet_address"`
	FullName           string             `json:"full_name"`
	CreationDate       string             `json:"creation_date"`
	PersonalData       PersonalData       `json:"PersonalData"`
	CurrentMedications CurrentMedications `json:"CurrentMedications"`
	PastMedicalHistory PastMedicalHistory `json:"PastMedicalHistory"`
	PersonalHistory    PersonalHistory    `json:"PersonalHistory"`
	FamilyHistory      FamilyHistory      `json:"FamilyHistory"`
	SystemsReview      SystemsReview      `json:"SystemsReview"`
	SubstanceUse       SubstanceUse       `json:"SubstanceUse"`
}

func (b *Block) generateHash() {
	// get string val of the Data
	bytes, _ := json.Marshal(b.Data)
	// concatenate the dataset
	data := string(b.Position) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, transaction Transaction) *Block {
	block := &Block{}
	block.Position = prevBlock.Position + 1
	block.Timestamp = time.Now().String()
	block.Data = transaction
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

// Blockchain is an ordered list of blocks
type Blockchain struct {
	blocks []*Block
}

// BlockChain is a global variable that'll return the mutated Blockchain struct
var BlockChain *Blockchain

// AddBlock adds a Block to a Blockchain
func (bc *Blockchain) AddBlock(data Transaction) {
	// get previous block
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// create new block
	block := CreateBlock(prevBlock, data)
	// validate integrity of blocks
	if validBlock(block, prevBlock) {
		// TODO: and if role permission okay
		bc.blocks = append(bc.blocks, block)
	}
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, Transaction{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func validBlock(block, prevBlock *Block) bool {
	// Confirm the hashes
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	// confirm the block's hash is valid
	if !block.validateHash(block.Hash) {
		return false
	}
	// Check the position to confirm its been incremented
	if prevBlock.Position+1 != block.Position {
		return false
	}
	return true
}

func validRole(userRole string) bool {
	// TODO: get user by UserID
	log.Printf("userRole: %v", userRole)
	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

func GetBlockchain(app *fiber.Ctx) error {
	jbytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
	if err != nil {
		app.JSON(http.StatusInternalServerError)
		json.NewEncoder(app).Encode(err)
		return nil
	}
	// write JSON string
	io.WriteString(app, string(jbytes))
	return nil
}

//func WriteBlock(w http.ResponseWriter, r *http.Request) {
//	var transaction Transaction
//	if validRole(transaction.UserRole) {
//		// Handle error
//		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			log.Printf("could not write Block: %v", err)
//			w.Write([]byte("could not write block"))
//			return
//		}
//
//		// Create block
//		BlockChain.AddBlock(transaction)
//		resp, err := json.MarshalIndent(transaction, "", " ")
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			log.Printf("could not marshal payload: %v", err)
//			w.Write([]byte("could not write block"))
//			return
//		}
//
//		// Response
//		w.WriteHeader(http.StatusOK)
//		w.Write(resp)
//	}
//}

func WriteBlock(c *fiber.Ctx) error {
	var transaction Transaction
	if validRole(transaction.UserRole) {
		// Handle error
		if err := c.BodyParser(&transaction); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("could not write block")
		}

		// Create block
		BlockChain.AddBlock(transaction)
		resp, err := json.MarshalIndent(transaction, "", " ")
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("could not write block")
		}

		// Response
		return c.Status(http.StatusSeeOther).Send(resp)
	}

	// Handle invalid role
	return c.Status(http.StatusUnauthorized).SendString("invalid role")
}

func NewMedicalRecord(c *fiber.Ctx) error {
	var medicalRecord MedicalRecord
	//if err := json.NewDecoder(r.Body).Decode(&medicalRecord); err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	log.Printf("could not create: %v", err)
	//	w.Write([]byte("could not create new MedicalRecord"))
	//	return
	//}
	//// We'll create an ID, concatenating the isdb and publish date
	//// This isn't an efficient way but serves for this tutorial
	//h := md5.New()
	//io.WriteString(h, medicalRecord.FullName+medicalRecord.CreationDate)
	//medicalRecord.WalletAddress = fmt.Sprintf("%x", h.Sum(nil))
	//
	//// send back payload
	//resp, err := json.MarshalIndent(medicalRecord, "", " ")
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	log.Printf("could not marshal payload: %v", err)
	//	w.Write([]byte("could not save medicalRecord data"))
	//	return
	//}
	//w.WriteHeader(http.StatusOK)
	//w.Write(resp)

	if err := c.BodyParser(&medicalRecord); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("could not create new MedicalRecord ok ?")

	}

	// We'll create an ID, concatenating the isdb and publish date
	// This isn't an efficient way but serves for this tutorial
	h := md5.New()
	_, err := io.WriteString(h, medicalRecord.FullName+medicalRecord.CreationDate)
	if err != nil {
		return err
	}
	medicalRecord.WalletAddress = fmt.Sprintf("%x", h.Sum(nil))

	// send back payload
	resp, err := json.MarshalIndent(medicalRecord, "", " ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("could not save medicalRecord data")
	}

	return c.Status(http.StatusOK).Send(resp)

}

//func newMedicalRecordPrint(w http.ResponseWriter, r *http.Request) {
//	var medicalRecord MedicalRecord
//
//	// We'll create an ID, concatenating the isdb and publish date
//	// This isn't an efficient way but serves for this tutorial
//	h := md5.New()
//	io.WriteString(h, medicalRecord.FullName+medicalRecord.CreationDate)
//	medicalRecord.WalletAddress = fmt.Sprintf("%x", h.Sum(nil))
//
//	w.WriteHeader(http.StatusOK)
//	w.Write(resp)
//}

func InitializeBC() {
	// initialize the blockchain and store in var
	BlockChain = NewBlockchain()

	// TODO: new user
	// r.HandleFunc("/user", newMedicalRecord).Methods("GET")
	// r.HandleFunc("/user", newMedicalRecord).Methods("POST")

	// dump the state of the Blockchain to the console
	go func() {
		//for {
		for _, block := range BlockChain.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
		//}
	}()
	log.Println("Listening on port 3001")

}
