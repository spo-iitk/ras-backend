package student

// ============================================================
// FILE: student/student_fetch.go
//
// WHAT THIS FILE DOES:
// Adds 2 new admin API endpoints:
//
//   1. GET /admin/student/:sid/data
//      → Give it a student's numeric DB id (sid)
//      → Returns full student profile as JSON
//
//   2. GET /admin/student/:sid/resume
//      → Give it a student's numeric DB id (sid)
//      → Returns ALL documents uploaded by that student
//        (because we don't know the exact "type" string used for resume,
//         this returns everything and the caller filters by type)
//        OR if you know the type string, we filter by it.
//
// WHY :sid IS A NUMBER:
//   Look at admin.document.go line:
//     sid, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
//   The existing code parses :sid as a uint (number like 1, 2, 3...)
//   It is the database primary key ID, NOT the roll number string.
//   So our new endpoints must do the same.
//
// IMPORTANT ABOUT RESUME TYPE:
//   In student.document.go, the Type field comes from whatever
//   the frontend sends in the JSON body — there's no hardcoded
//   string like "resume" in the backend code we can see.
//   So getStudentResumeHandler returns ALL documents for that student.
//   The frontend/caller can then filter by type themselves.
//   (If you find the type string later, just uncomment the filter line)
// ============================================================

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ─────────────────────────────────────────────────────────────────────────────
// ENDPOINT 1: GET /admin/student/:sid/data
//
// Returns full profile of a student given their numeric DB id.
//
// Example:
//
//	GET /admin/student/5/data
//	→ returns the Student row where ID = 5
//
// Example JSON response:
//
//	{
//	  "ID": 5,
//	  "roll_no": "210101",
//	  "name": "Rahul Sharma",
//	  "iitk_email": "rahuls21@iitk.ac.in",
//	  "current_cpi": 8.5,
//	  "gender": "Male",
//	  ... all other fields from Student struct in model.go
//	}
//
// ─────────────────────────────────────────────────────────────────────────────
func getStudentDataHandler(ctx *gin.Context) {
	// Step 1: Read :sid from URL and convert it to a number
	// strconv.ParseUint converts the string "5" → uint64(5)
	// The existing code in admin.document.go does the exact same thing
	sid, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
	if err != nil {
		// If someone passes /admin/student/abc/data → bad request
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student id: " + ctx.Param("sid"),
		})
		return
	}

	// Step 2: Fetch the student from DB using their numeric ID
	// getStudentByID already exists in your codebase (used in admin.document.go)
	// It queries: SELECT * FROM students WHERE id = sid
	var student Student
	err = getStudentByID(ctx, &student, uint(sid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Step 3: Return the student as JSON
	ctx.JSON(http.StatusOK, student)
}

// ─────────────────────────────────────────────────────────────────────────────
// ENDPOINT 2: GET /admin/student/:sid/resume
//
// Returns all documents uploaded by a student given their numeric DB id.
// The "resume" is one of these documents — its Type field tells you what it is.
//
// Example:
//
//	GET /admin/student/5/resume
//	→ returns all StudentDocument rows where student_id = 5
//
// Example JSON response:
//
//	[
//	  {
//	    "ID": 12,
//	    "StudentID": 5,
//	    "type": "resume",        ← whatever string the frontend used
//	    "path": "https://...",   ← THIS IS THE ACTUAL FILE URL
//	    "verified": true,
//	    "action_taken_by": "admin@iitk.ac.in"
//	  },
//	  {
//	    "ID": 13,
//	    "StudentID": 5,
//	    "type": "transcript",
//	    "path": "https://...",
//	    ...
//	  }
//	]
//
// NOTE: We reuse getDocumentsByStudentID from db.document.go —
//
//	the exact same function the existing getDocumentHandler uses!
//	We are just calling it from a new route.
//
// ─────────────────────────────────────────────────────────────────────────────
func getStudentResumeHandler(ctx *gin.Context) {
	// Step 1: Read :sid from URL and convert to number
	// Same as getDocumentHandler in admin.document.go does
	sid, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student id: " + ctx.Param("sid"),
		})
		return
	}

	// Step 2: Fetch all documents for this student
	// getDocumentsByStudentID is defined in db.document.go:
	//   func getDocumentsByStudentID(ctx, documents, studentID) error
	//   it runs: SELECT * FROM student_documents WHERE student_id = sid
	var documents []StudentDocument
	err = getDocumentsByStudentID(ctx, &documents, uint(sid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Step 3: Check if any documents exist
	if len(documents) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No documents found for student id: " + ctx.Param("sid"),
		})
		return
	}

	// Step 4: Return all documents as JSON
	// The caller can look at the "type" field to find the resume
	ctx.JSON(http.StatusOK, documents)

	// ── OPTIONAL ──────────────────────────────────────────────
	// If you find out the exact type string used for resumes
	// (e.g. "resume" or "Resume" or "RESUME"), you can filter like this:
	//
	// var resumeDocs []StudentDocument
	// for _, doc := range documents {
	//     if doc.Type == "resume" {   // ← replace with actual type string
	//         resumeDocs = append(resumeDocs, doc)
	//     }
	// }
	// if len(resumeDocs) == 0 {
	//     ctx.JSON(http.StatusNotFound, gin.H{"error": "No resume found"})
	//     return
	// }
	// ctx.JSON(http.StatusOK, resumeDocs)
	// ──────────────────────────────────────────────────────────
}
