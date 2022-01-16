/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */


func firstBadVersion(n int) int {
    l, r := 1, n
    
    for l < r {
        mid := l+(r-l)/2
        isBad := isBadVersion(mid)
        if !isBad {
            l = mid+1
        } else {
            r = mid
        }
    }
    if isBadVersion(r) {
        return r
    }
    return 0
}