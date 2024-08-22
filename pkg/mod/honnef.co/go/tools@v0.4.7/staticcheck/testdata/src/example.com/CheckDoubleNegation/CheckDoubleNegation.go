package pkg

func fn(b1, b2 bool) {
	if !!b1 { //@ diag(`negating a boolean twice`)
		println()
	}

	if b1 && !!b2 { //@ diag(`negating a boolean twice`)
		println()
	}

	if !(!b1) { //@ diag(`negating a boolean twice`)
		println()
	}

	if !b1 {
		println()
	}

	if !b1 && !b2 {
		println()
	}
}
