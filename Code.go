package main

import (
	"fmt"
)

const NMAX = 100
const NMAX_BALASAN = 10

type User struct {
	Username string
	Password string
}

type Reply struct {
	User    string
	Content string
}

type Question struct {
	Pertanyaan string
	Content    string
	Tag        string
	replies    [NMAX_BALASAN]Reply
	mostAsked  int
	NumReplies int
}

type Forum struct {
	Questions   [NMAX]Question
	NumQuestion int
}

var users [NMAX]User
var forum Forum
var username string

func addComment() {
	fmt.Println("=== Balas Pertanyaan ===")
	fmt.Println("Pilih pertanyaan yang ingin dibalas:")
	for i, question := range forum.Questions {
		fmt.Printf("%d. %s (mostAsked: %d, Balasan: %d)\n", i+1, question.Pertanyaan, question.mostAsked, question.NumReplies)
	}
	fmt.Print("Pilih nomor: ")

	var choice int
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(forum.Questions) {
		fmt.Println("Silakan posting balasan Anda:")
		var replyContent string
		replyContent = readMultilineInput()

		reply := Reply{
			User:    username,
			Content: replyContent,
		}

		forum.Questions[choice-1].replies[forum.Questions[choice-1].NumReplies] = reply
		forum.Questions[choice-1].NumReplies++
		fmt.Println("Balasan telah diposting. Kembali ke menu utama.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func registerAccount() {
	fmt.Println("=== Pendaftaran Pasien ===")
	fmt.Print("Nama Pasien: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Masukkan kata sandi: ")
	var password string
	fmt.Scanln(&password)

	var i int
	for i = 0; i < NMAX && users[i].Username != ""; i++ {
	}

	if i == NMAX {
		fmt.Println("Kapasitas Pasien telah mencapai batas.")
		return
	}

	users[i].Username = username
	users[i].Password = password

	fmt.Println("Pendaftaran berhasil!")
}

func postQuestion() {
	fmt.Println("=== Forum ===")
	fmt.Println("Silakan posting pertanyaan Anda:")
	var questionTitle, questionContent string
	fmt.Print("Judul Pertanyaan: ")
	fmt.Scanln(&questionTitle)
	fmt.Print("Isi Pertanyaan: ")
	fmt.Scanln(&questionContent)
	fmt.Print("Tag Pertanyaan: ")
	var questionTag string
	fmt.Scanln(&questionTag)

	question := Question{
		Pertanyaan: questionTitle,
		Content:    questionContent,
		Tag:        questionTag,
		mostAsked:  0,
	}

	forum.Questions[forum.NumQuestion] = question
	forum.NumQuestion++

	fmt.Println("Pertanyaan telah diposting. Kembali ke menu utama.")
}

func readMultilineInput() string {
	const maxLines = 100
	lines := [maxLines]string{}
	var isBreak bool = false
	var count int = 0
	for !isBreak && count < maxLines {
		var line string
		fmt.Scanln(&line)
		if line == "" {
			isBreak = true
		} else {
			lines[count] = line
			count++
		}
	}

	var result string
	for i := 0; i < count; i++ {
		if i > 0 {
			result += " "
		}
		result += lines[i]
	}
	return result
}

func joinLines(lines []string, separator string) string {
	result := ""
	for i, line := range lines {
		result += line
		if i < len(lines)-1 {
			result += separator
		}
	}
	return result
}

func showForum() {
	fmt.Println("=== Forum ===")
	fmt.Println("Pertanyaan yang telah diposting:")

	if len(forum.Questions) == 0 {
		fmt.Println("Belum ada pertanyaan yang diposting.")
		return
	}

	for i := 0; i < forum.NumQuestion; i++ {
		fmt.Printf("%d. %s (mostAsked: %d, Balasan: %d)\n", i+1, forum.Questions[i].Pertanyaan, forum.Questions[i].mostAsked, forum.Questions[i].NumReplies)
		fmt.Println("   Isi Pertanyaan:")
		fmt.Println("   ", forum.Questions[i].Content)
		fmt.Println("   Tag:", forum.Questions[i].Tag)
		fmt.Println("   Balasan:")
		for j := 0; j < forum.Questions[i].NumReplies; j++ {
			fmt.Println("   ", forum.Questions[i].replies[j].Content)
		}
		fmt.Println()
	}

	fmt.Println("Pilih nomor pertanyaan untuk melihat detail atau tekan 0 untuk kembali.")
	fmt.Print("Pilih nomor: ")

	var choice int
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(forum.Questions) {
		question := forum.Questions[choice-1]
		fmt.Printf("=== Detail Pertanyaan ===\nJudul: %s\nmostAsked: %d\nBalasan: %d\nIsi Pertanyaan:\n%s\n", question.Pertanyaan, question.mostAsked, question.NumReplies, question.Content)
	} else if choice != 0 {
		fmt.Println("Pilihan tidak valid.")
	}
}

func formatText(text string, indent int) string {
	indentation := ""
	for i := 0; i < indent; i++ {
		indentation += " "
	}
	text = text + indentation
	return indentation + text
}

func showTopQuestions() {
	questions := forum.Questions
	n := len(questions)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if questions[j].mostAsked < questions[j+1].mostAsked {
				temp := questions[j]
				questions[j] = questions[j+1]
				questions[j+1] = temp
			}
		}
	}

	fmt.Println("=== Jumlah Pertanyaan berdasarkan Kategori Terbanyak ===")
	for i, question := range questions {
		fmt.Printf("%d. %s (Jumlah orang yang bertanya: %d)\n", i+1, question.Tag, question.mostAsked)
	}
}

func showTopReplies() {
	questions := forum.Questions
	n := len(questions)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if questions[j].NumReplies < questions[j+1].NumReplies {
				temp := questions[j]
				questions[j] = questions[j+1]
				questions[j+1] = temp
			}
		}
	}

	fmt.Println("=== Pertanyaan dengan Balasan Terbanyak ===")
	for i, question := range questions {
		fmt.Printf("%d. %s (Balasan: %d)\n", i+1, question.Pertanyaan, question.NumReplies)
	}
}

func login() {
	fmt.Println("=== Selamat Datang! ===")
	fmt.Print("Nama Pasien: ")
	fmt.Scanln(&username)

	fmt.Print("Masukkan kata sandi: ")
	var password string
	fmt.Scanln(&password)

	for _, user := range users {
		if user.Username == username && user.Password == password {
			fmt.Println("Login berhasil!")
			for {
				fmt.Println("\nPilihan:")
				fmt.Println("1. Posting Pertanyaan di Forum")
				fmt.Println("2. Lihat Pertanyaan yang Telah Diposting")
				fmt.Println("3. Tambah Balasan ke Pertanyaan")
				fmt.Println("4. Lihat Pertanyaan dengan Kategori terbanyak")
				fmt.Println("5. Lihat Pertanyaan dengan Balasan Terbanyak")
				fmt.Println("6. Keluar")
				fmt.Print("Pilih nomor menu: ")

				var choice int
				fmt.Scanln(&choice)

				if choice == 1 {
					postQuestion()
				} else if choice == 2 {
					showForum()
				} else if choice == 3 {
					addComment()
				} else if choice == 4 {
					showTopQuestions()
				} else if choice == 5 {
					showTopReplies()
				} else if choice == 6 {
					fmt.Println("Terima kasih!")
					return
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		}
	}

	fmt.Println("Nama Pasien atau kata sandi salah.")
}

func main() {
	addDataDummyForum()
	forum.NumQuestion += 3
	
	fmt.Println(len(forum.Questions))
	

	fmt.Println("	=== Aplikasi Konsultasi Kesehatan ===")
	fmt.Println("***                  Created by                 ***")
    fmt.Println("***               Ade Ikmal Maulana             ***")
    fmt.Println("***                  M. Naufal 		        ***")
    fmt.Println("***          Algoritma Pemrograman 2023         ***")


	for {
		fmt.Println("\nPilihan:")
		fmt.Println("1. Daftar Pasien Baru")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Forum")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih nomor menu: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			registerAccount()
		} else if choice == 2 {
			login()
		} else if choice == 3 {
			showForum()
		} else if choice == 4 {
			fmt.Println("Terima kasih!")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addDataDummyForum() {
	forum.Questions = [NMAX]Question{
		{
			Pertanyaan: "Apakah merokok dapat meningkatkan resiko penyakit jantung?",
			Content:    "Setiap hari, ayah saya menghabiskan satu kemasan rokok, saya khawatir pada keadaan ayah saya",
			Tag:        "Penyakit_Dalam",
			mostAsked:  0,
			NumReplies: 0,
		},
		{
			Pertanyaan: "Apakah_menggunakan_alkohol_untuk_mengobati_luka_luar_aman?",
			Content:    "Saya_telah_menggunakan_alkohol_dalam_mengobati_luka_saya_selama_bertahun-tahun,_apakah_hal_tersebut_aman?",
			Tag:        "Penyakit_Luar",
			mostAsked:  0,
			NumReplies: 0,
		},
		{
			Pertanyaan: "Apakah_menjaga_mental_health_penting?",
			Content:    "Saya_telah_bekerja_pada_perusahaan_yang_gajinya_cukup_besar.Akan_tetapi,_saya_sering_stress_dalam_menjalaninya_sampai_kesehatan_saya_terganggu._Haruskah_saya_resign_dan_menjalani_pola_hidup_sehat?",
			Tag:        "Psikologi",
			mostAsked:  0,
			NumReplies: 0,
		},
	}
}