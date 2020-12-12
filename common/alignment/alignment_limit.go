/*
 * @author: Haoyuan Liu
 * @date: 2020/11/30
 */

package alignment

type Limiter interface {
	AlignmentLimit(alignment Alignment) bool
}
