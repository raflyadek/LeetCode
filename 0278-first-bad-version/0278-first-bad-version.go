/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    res := n
    left, right := 0, n 
    for left <= right {
        mid := (left+right)/2
        bad := isBadVersion(mid)
        if bad {
            if mid < res {
                res = mid
            }
            right = mid - 1
        } else {
            left = mid + 1
        }
    }    
    return res
}