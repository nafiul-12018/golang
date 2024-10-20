


func parseBoolExpr(expression string) bool {
     var op []rune
     var val []rune
     for _, ch := range expression {
        if ch == ')' {
            t := 0
            f := 0
            opIdx := len(op)-1
            idx := len(val) -1
            for val[idx] != '(' {
                if val[idx] == 'f' {
                    f = 1
                }

                if val[idx] == 't' {
                    t = 1
                }
                idx--;  
            }
            val = val[:idx]
            //fmt.Println(t,f)

            if op[opIdx] == '!'  {
                if t == 1 {
                    val = append(val, 'f')
                } else {
                    val = append(val, 't')
                }
            } else {
                if t == 0 {
                    val = append(val, 'f')
                } else if f == 0 {
                    val = append(val, 't')
                } else if op[opIdx] == '|' {
                    val = append(val, 't')
                } else {
                     val = append(val, 'f')
                }
            }
            op = op[:opIdx]
            continue
        }
        
        if ch == ',' {
            continue
        }

        if ch == 't' || ch == 'f' || ch =='(' {
            val = append(val, ch)
        } else {
            op = append(op, ch)
        }
     }
     return val[0] == 't'
}
