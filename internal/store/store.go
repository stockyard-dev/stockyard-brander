package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Asset struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Format string `json:"format"`
	URL string `json:"url"`
	Version string `json:"version"`
	Status string `json:"status"`
	Notes string `json:"notes"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"brander.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS assets(id TEXT PRIMARY KEY,name TEXT NOT NULL,type TEXT DEFAULT 'logo',format TEXT DEFAULT '',url TEXT DEFAULT '',version TEXT DEFAULT '1.0',status TEXT DEFAULT 'active',notes TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Asset)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO assets(id,name,type,format,url,version,status,notes,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Type,e.Format,e.URL,e.Version,e.Status,e.Notes,e.CreatedAt);return err}
func(d *DB)Get(id string)*Asset{var e Asset;if d.db.QueryRow(`SELECT id,name,type,format,url,version,status,notes,created_at FROM assets WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Type,&e.Format,&e.URL,&e.Version,&e.Status,&e.Notes,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Asset{rows,_:=d.db.Query(`SELECT id,name,type,format,url,version,status,notes,created_at FROM assets ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Asset;for rows.Next(){var e Asset;rows.Scan(&e.ID,&e.Name,&e.Type,&e.Format,&e.URL,&e.Version,&e.Status,&e.Notes,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Asset)error{_,err:=d.db.Exec(`UPDATE assets SET name=?,type=?,format=?,url=?,version=?,status=?,notes=? WHERE id=?`,e.Name,e.Type,e.Format,e.URL,e.Version,e.Status,e.Notes,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM assets WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM assets`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Asset{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (name LIKE ?)"
        args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["type"];ok&&v!=""{where+=" AND type=?";args=append(args,v)}
    if v,ok:=filters["status"];ok&&v!=""{where+=" AND status=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,name,type,format,url,version,status,notes,created_at FROM assets WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Asset;for rows.Next(){var e Asset;rows.Scan(&e.ID,&e.Name,&e.Type,&e.Format,&e.URL,&e.Version,&e.Status,&e.Notes,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    rows,_:=d.db.Query(`SELECT status,COUNT(*) FROM assets GROUP BY status`)
    if rows!=nil{defer rows.Close();by:=map[string]int{};for rows.Next(){var s string;var c int;rows.Scan(&s,&c);by[s]=c};m["by_status"]=by}
    return m
}
