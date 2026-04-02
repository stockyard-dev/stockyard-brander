package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Asset struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	URL string `json:"url"`
	Description string `json:"description"`
	FileType string `json:"file_type"`
	Version string `json:"version"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"brander.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS assets(id TEXT PRIMARY KEY,name TEXT NOT NULL,category TEXT DEFAULT '',url TEXT DEFAULT '',description TEXT DEFAULT '',file_type TEXT DEFAULT '',version TEXT DEFAULT '1.0',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Asset)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO assets(id,name,category,url,description,file_type,version,created_at)VALUES(?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Category,e.URL,e.Description,e.FileType,e.Version,e.CreatedAt);return err}
func(d *DB)Get(id string)*Asset{var e Asset;if d.db.QueryRow(`SELECT id,name,category,url,description,file_type,version,created_at FROM assets WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Category,&e.URL,&e.Description,&e.FileType,&e.Version,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Asset{rows,_:=d.db.Query(`SELECT id,name,category,url,description,file_type,version,created_at FROM assets ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Asset;for rows.Next(){var e Asset;rows.Scan(&e.ID,&e.Name,&e.Category,&e.URL,&e.Description,&e.FileType,&e.Version,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM assets WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM assets`).Scan(&n);return n}
