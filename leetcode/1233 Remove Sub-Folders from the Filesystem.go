
type Folder struct {
    IsEnd bool
    Child map[string]*Folder
}

func insertTrie (folder *Folder, data string) {
    words := strings.Split(data, "/")
    for _,w := range words {
        if len(w) == 0{
            continue
        }
            
        ch, exists := folder.Child[w]

        if exists == false {
            ch = &Folder{
                Child : make(map[string]*Folder),
            }
            folder.Child[w] = ch
        }
        if ch.IsEnd == true {
            return
        }
        folder = ch
    }
    folder.IsEnd = true
}

func dfs(folder *Folder, tmp []string, ans *[]string) {
    if folder.IsEnd == true {
        tmpstr := "/" + strings.Join(tmp, "/")
        *ans = append(*ans,tmpstr)
        return 
    }
    for key, val := range folder.Child {
        dfs(val, append(tmp, key), ans)
    }
}

func removeSubfolders(folder []string) []string {

    root := &Folder{
        Child : make(map[string]*Folder),
    }

    for _, data := range folder {
        //fmt.Println(data)
        insertTrie(root, data)
    }

    //fmt.Println(root)
    tmp := []string{}
    ans := []string{}
    dfs(root, tmp, &ans)

    return ans
}
